package main

import (
	"os"
)

type Cfg struct {
	AccessKey string
	SecretKey string
}

func Config() *Cfg {
	cfg := &Cfg{}

	cfg.AccessKey = os.Getenv("PWNIE_ACCESS_KEY")
	cfg.SecretKey = os.Getenv("PWNIE_SECRET_KEY")

	return cfg
}
