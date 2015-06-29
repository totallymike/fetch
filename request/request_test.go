package request

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	os.Setenv("AUTH_REGION", "us-east-1")
	m.Run()
}
