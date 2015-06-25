package main

import (
	"testing"
	"time"
)

func newRequest() (signedRequest *SignedRequest) {
	signedRequest, _ = NewSignedRequest(
		"GET",
		"http://www.example.com/v1/network-hosts?foo=bar&baz=foo",
	)
	return
}

func TestCanonicalURI(t *testing.T) {
	t.Parallel()
	req := newRequest()

	if req.CanonicalURI != "/v1/network-hosts" {
		t.Errorf("%s != %s\n", req.CanonicalURI, "/v1/network-hosts")
	}
}

func TestCanonicalQueryString(t *testing.T) {
	t.Parallel()
	req := newRequest()

	expected := "baz=foo&\nfoo=bar"
	if req.CanonicalQueryString() != expected {
		t.Errorf("%s != %s\n", req.CanonicalQueryString(), expected)
	}
}

func TestCanonicalHeaders(t *testing.T) {
	t.Parallel()
	req := newRequest()

	expected := "content-type:application/vnd.api+json\n" +
		"host:www.example.com\n" +
		"x-amz-date:" + time.Now().Format("20060102T150405Z")

	if req.CanonicalHeaders() != expected {
		t.Errorf("%s != %s\n", req.CanonicalHeaders(), expected)
	}
}

func TestCanonicalHeadersWithSpaces(t *testing.T) {
	t.Parallel()
	req := newRequest()
	req.AddHeader("foo", `"   oh  yeah"`)
	req.AddHeader("bar", "   oh     yeah")

	expected := "bar:oh yeah\ncontent-type:application/vnd.api+json\n" +
		"foo:\"   oh  yeah\"\nhost:www.example.com\n" +
		"x-amz-date:" + time.Now().Format("20060102T150405Z")


	if req.CanonicalHeaders() != expected {
		t.Errorf("%s != %s\n", req.CanonicalHeaders(), expected)
	}
}


/*
func TestCanonicalRequest(t *testing.T) {
	t.Skip()
	t.Parallel()

	req := newRequest()

	expected := "GET\n/v1/network-hosts\nbaz=foo&\nfoo=bar\n" +
		"content-type:application/vnd.api+json\nhost:www.example.com"
}
*/
