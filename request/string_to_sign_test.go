package request

import (
	"testing"
	"time"
)

func TestStringToSign(t *testing.T) {
	t.Parallel()
	t.Log("boogers")
	now := time.Now().UTC()
	req := newRequest()

	expected := "AWS4-HMAC-SHA256\n" +
		now.Format("20060102T150405Z") + "\n" +
		now.Format("20060102") + "/us-east-1/www/aws4_request\n" +
		req.HashedCanonicalRequest("")

	actual := req.StringToSign("")

	if expected != actual {
		t.Errorf("%s != %s\n", actual, expected)
	}
}
