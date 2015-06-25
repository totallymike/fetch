package main

import (
	"testing"
	"os"
	"time"
	"fmt"
)

func setupEnvironment() {
	access_key := "my_access_key"
	secret_key := "foobar"
	region := "us-east-1"

	os.Setenv("PWNIE_ACCESS_KEY", access_key)
	os.Setenv("PWNIE_SECRET_KEY", secret_key)
	os.Setenv("PWNIE_REGION", region)
}

func TestAddAuthorizationHeader(t *testing.T) {
	t.Parallel()
	setupEnvironment()

	req := newRequest()
	payload := ""
	req.AddAuthorizationHeader(payload)

	if _, ok := req.Header()["Authorization"]; !ok {
		t.Errorf("Authorization header not found")
	}
}

func TestAuthorizationHeader(t *testing.T) {
	t.Parallel()

	setupEnvironment()

	payload := ""
	req := newRequest()
	now := time.Now()

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := formatShortDate(now) + "/us-east-1/www/aws4_request"
	signedHeaders := req.SignedHeaders()
	signature := req.Signature(payload)

	expected := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		os.Getenv("PWNIE_ACCESS_KEY"),
		credentialScope,
		signedHeaders,
		signature,
	)

	actual := req.AuthorizationHeader(payload)
	if expected != actual {
		t.Errorf("%s != %s\n", actual, expected)
	}
}
