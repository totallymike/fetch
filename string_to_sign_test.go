package main

import (
	"testing"
	"time"
)

func TestStringToSign(t *testing.T) {
	t.Parallel()
	req := newRequest()

	req.Config.Region = "us-east-1"

	expected := "AWS4-HMAC-SHA256\n" +
		time.Now().Format("20060102T150405Z") + "\n" +
		time.Now().Format("20060102") + "/us-east-1/example.com/aws4_request\n" +
		req.HashedCanonicalRequest("")

	actual := req.StringToSign("")

	if expected != actual {
		t.Errorf("%s != %s\n", actual, expected)
	}
}
