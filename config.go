package main

import (
	"os"
)

type Cfg struct {
	AccessKey string
	SecretKey string
	Region string
}

func Config() *Cfg {
	cfg := &Cfg{}

	cfg.AccessKey = os.Getenv("PWNIE_ACCESS_KEY")
	cfg.SecretKey = os.Getenv("PWNIE_SECRET_KEY")
	if cfg.Region = os.Getenv("PWNIE_REGION"); cfg.Region == "" {
		cfg.Region = "us-east-1"
	}

	return cfg
}
