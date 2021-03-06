package request

import (
	"testing"
	"os"
	"crypto/hmac"
	"crypto/sha256"
	"time"
	"encoding/hex"
	"github.com/spf13/viper"
)

func TestSignature(t *testing.T) {
	t.Parallel()
	t.Log("o")
	req := newRequest()

	mac := hmac.New(sha256.New, req.DerivedSigningKey())
	mac.Write([]byte(req.StringToSign("")))
	expected := hex.EncodeToString(mac.Sum(nil))

	actual := req.Signature("")

	if expected != actual {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestDerivededSigningKey(t *testing.T) {
	t.Parallel()

	os.Setenv("PWNIE_SECRET_KEY", "foobar")
	os.Setenv("PWNIE_REGION", "us-east-1")
	req := newRequest()

	mac := hmac.New(sha256.New, []byte("AWS4foobar"))
	mac.Write([]byte(formatShortDate(time.Now().UTC())))
	kDate := mac.Sum(nil)

	mac = hmac.New(sha256.New, kDate)
	mac.Write([]byte(viper.GetString("region")))
	kRegion := mac.Sum(nil)

	mac = hmac.New(sha256.New, kRegion)
	mac.Write([]byte("www"))
	kService := mac.Sum(nil)

	mac = hmac.New(sha256.New, kService)
	mac.Write([]byte("aws4_request"))
	kSigning := mac.Sum(nil)

	expected := kSigning
	actual := req.DerivedSigningKey()

	if !hmac.Equal(expected, actual) {
		t.Errorf("\n%v\n != \n%v", actual, expected)
	}
}
