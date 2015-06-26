package main

import (
	"os"
	"flag"
)

type Cfg struct {
	AccessKey string
	SecretKey string
	AllowInsecureSsl bool
	Region string
}

func Config() *Cfg {
	cfg := &Cfg{}

	access_key := flag.String("access_key", "", "your access key")
	secret_key := flag.String("secret_key", "", "your secret key")
	insecure_ssl := flag.Bool("insecure", false, "allow bad ssl certs")

	flag.Parse()

	cfg.AccessKey = *access_key
	cfg.SecretKey = *secret_key
	cfg.AllowInsecureSsl = *insecure_ssl

	if cfg.Region = os.Getenv("PWNIE_REGION"); cfg.Region == "" {
		cfg.Region = "us-east-1"
	}

	return cfg
}
