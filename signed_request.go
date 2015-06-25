package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

type SignedRequest struct {
	CanonicalURI string
	client http.Client
	request http.Request
}

func NewSignedRequest(method string, url string) (signedRequest *SignedRequest, err error) {
	signedRequest = &SignedRequest{}

	req, err := http.NewRequest("GET",
		url,
		nil,
	)

	signedRequest.request = *req

	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.Header.Add("Host", req.URL.Host)
	req.Header.Add("X-AMZ-DATE", time.Now().Format("20060102T150405Z"))
	signedRequest.CanonicalURI = req.URL.Path
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

func (req *SignedRequest) Header() http.Header {
	return req.request.Header
}

func (req *SignedRequest) AddHeader(name string, value string) {
	req.request.Header.Add(name, value)
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
		if strings.HasPrefix(valStr, `"`) && strings.HasSuffix(valStr, `"`) {
			valStr = valStr
		} else {
			valStr = whitespaceTrimPattern.ReplaceAllString(valStr, " ")
			valStr = strings.TrimSpace(valStr)
		}

		canonicizedHeaders[strings.ToLower(k)] = []string{valStr}
	}
	return getCanonicalForm(canonicizedHeaders, ":", "")
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
