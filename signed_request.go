package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
	"net/url"
)

type SignedRequest struct {
	CanonicalURI string
	client       *http.Client
	request      http.Request
	service      string
	Config       *Cfg
	TimeOfRequest time.Time
}

func NewSignedRequest(method string, requestUrl string) (
	signedRequest *SignedRequest, err error,
) {
	signedRequest = &SignedRequest{}
	signedRequest.Config = Config()

	parsedUrl, err := url.Parse(requestUrl)

	req, err := http.NewRequest(method,
		parsedUrl.String(),
		nil,
	)

	signedRequest.request = *req
	signedRequest.TimeOfRequest = time.Now().UTC()

	service := strings.SplitN(req.URL.Host, ".", 2)[0]
	if strings.ContainsRune(service, ':') {
		service = strings.SplitN(service, ":", 2)[0]
	}

	signedRequest.service = service

	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.Header.Add("Host", req.URL.Host)
	req.Header.Add("X-AMZ-DATE", signedRequest.TimeOfRequest.Format("20060102T150405Z"))
	signedRequest.CanonicalURI = req.URL.Path
	return
}

func (req *SignedRequest) Perform(payload string) (response *http.Response, err error) {
	if req.client == nil {
		req.client = &http.Client{}
	}

	req.AddAuthorizationHeader(payload)

	response, err = req.client.Do(&req.request)
	return
}

func (req *SignedRequest) CanonicalQueryString() string {
	return getCanonicalForm(req.request.URL.Query(), "=", "&")
}

func (req *SignedRequest) CanonicalHeaders() string {
	return getCanonicalHeaders(req.Header())
}

func (req *SignedRequest) SignedHeaders() string {
	headers := req.Header()
	signedHeaders := make([]string, len(headers))

	i := 0
	for k := range headers {
		signedHeaders[i] = strings.ToLower(k)
		i += 1
	}

	sort.Strings(signedHeaders)

	return strings.Join(signedHeaders, ";")
}

func (req *SignedRequest) SignedPayload(payload string) string {
	return signPayload(payload)
}

func (req *SignedRequest) CanonicalRequest(payload string) string {
	return strings.Join(
		[]string{
			req.request.Method,
			req.CanonicalURI,
			req.CanonicalQueryString(),
			req.CanonicalHeaders(),
			req.SignedHeaders(),
			req.SignedPayload(payload),
		},
		"\n",
	)
}

func (req *SignedRequest) AddAuthorizationHeader(payload string) {
	req.Header().Add("Authorization", req.AuthorizationHeader(payload))
}

func (req *SignedRequest) AuthorizationHeader(payload string) string {
	nowString := formatShortDate(req.TimeOfRequest)
	region := req.Config.Region
	service := strings.SplitN(req.request.URL.Host, ".", 2)[0]

	credentialScope := fmt.Sprintf("%s/%s/%s/aws4_request",
		nowString,
		region,
		service)

	return fmt.Sprintf(
		"AWS4-HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		req.Config.AccessKey,
		credentialScope,
		req.SignedHeaders(),
		req.Signature(payload))
}

func (req *SignedRequest) HashedCanonicalRequest(payload string) string {
	return signPayload(req.CanonicalRequest(payload))
}

func (req *SignedRequest) StringToSign(payload string) string {
	now := req.TimeOfRequest
	return "AWS4-HMAC-SHA256\n" +
		formatLongDate(now) + "\n" +
		formatShortDate(now) + "/" + req.Config.Region +
		"/" + req.service + "/aws4_request\n" +
		req.HashedCanonicalRequest(payload)
}

func (req *SignedRequest) DerivedSigningKey() []byte {
	now := req.TimeOfRequest
	secretKey := "AWS4" + req.Config.SecretKey
	region := req.Config.Region

	kDate := hmacHash([]byte(secretKey), []byte(formatShortDate(now)))
	kRegion := hmacHash(kDate, []byte(region))
	kService := hmacHash(kRegion, []byte(req.service))

	return hmacHash(kService, []byte("aws4_request"))
}

func (req *SignedRequest) Signature(payload string) string {
	stringToSign := req.StringToSign(payload)
	hashedSignature := hmacHash(req.DerivedSigningKey(), []byte(stringToSign))
	return hex.EncodeToString(hashedSignature)
}

func (req *SignedRequest) Header() http.Header {
	return req.request.Header
}

func (req *SignedRequest) AddHeader(name string, value string) {
	req.request.Header.Add(name, value)
}

func formatShortDate(date time.Time) string {
	return date.Format("20060102")
}

func formatLongDate(date time.Time) string {
	return date.Format("20060102T150405Z")
}

func hmacHash(key, content []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(content)
	return mac.Sum(nil)
}

func signPayload(payload string) string {
	hash := sha256.New()
	io.WriteString(hash, payload)
	return hex.EncodeToString(hash.Sum(nil))
}

func getCanonicalHeaders(data map[string][]string) string {
	whitespaceTrimPattern := regexp.MustCompile(`\s{2,}`)

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	canonicizedHeaders := make(map[string][]string)

	for _, k := range keys {
		val := data[k]

		valStr := strings.Join(val, "")

		if !(strings.HasPrefix(valStr, `"`) && strings.HasSuffix(valStr, `"`)) {
			valStr = whitespaceTrimPattern.ReplaceAllString(valStr, " ")
			valStr = strings.TrimSpace(valStr)
		}

		canonicizedHeaders[strings.ToLower(k)] = []string{valStr}
	}
	return getCanonicalForm(canonicizedHeaders, ":", "") + "\n"
}

func getCanonicalForm(data map[string][]string, kvJoin string, entryJoin string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var queryStrings []string
	for _, k := range keys {
		val := strings.Join(data[k], "")
		queryStrings = append(queryStrings, fmt.Sprintf("%s%s%s", k, kvJoin, val))
	}

	return strings.Join(queryStrings, fmt.Sprintf("%s\n", entryJoin))
}
