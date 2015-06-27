package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("AUTH_ACCESS_KEY", "Foo")
	os.Setenv("AUTH_SECRET_KEY", "Bar")
	os.Exit(m.Run())
}

func failTest(t *testing.T, expected interface{}, actual interface{}) {
	t.Log("FAIL", expected, "does not match", actual)
}

func TestAccessKey(t *testing.T) {
	cfg := Config()

	if cfg.AccessKey != "Foo" {
		failTest(t, "Foo", cfg.AccessKey)
		t.Fail()
	}
}

func TestSecretKey(t *testing.T) {
	cfg := Config()

	if cfg.SecretKey != "Bar" {
		failTest(t, "Bar", cfg.SecretKey)
		t.Fail()
	}
}
