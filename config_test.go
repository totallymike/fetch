package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("PWNIE_ACCESS_KEY", "Foo")
	os.Setenv("PWNIE_SECRET_KEY", "Bar")
	os.Exit(m.Run())
}

func failTest(t *testing.T, expected interface{}, actual interface{}) {
	t.Error("FAIL", expected, "does not match", actual)
}

func TestAccessKey(t *testing.T) {
	t.Parallel()
	cfg := Config()

	if cfg.AccessKey != "Foo" {
		failTest(t, "Foo", cfg.AccessKey)
	}
}

func TestSecretKey(t *testing.T) {
	t.Parallel()
	cfg := Config()

	if cfg.SecretKey != "Bar" {
		failTest(t, "Bar", cfg.AccessKey)
	}
}
