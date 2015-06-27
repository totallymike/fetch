package main

import (
	"github.com/spf13/viper"
	"github.com/totallymike/aws-authenticated-request-thing/commands"
)

type Cfg struct {
	AccessKey string
	SecretKey string
	AllowInsecureSsl bool
	Region string
}

var config *Cfg

func Config() *Cfg {
	commands.Execute()

	cfg := &Cfg{}

	insecure_ssl := false

	cfg.AccessKey = viper.GetString("access_key")
	cfg.SecretKey = viper.GetString("secret_key")
	cfg.AllowInsecureSsl = insecure_ssl

	cfg.Region = viper.GetString("Region")

	return cfg
}
