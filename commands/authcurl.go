package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

var authCmdV *cobra.Command
var AuthCurlCmd = &cobra.Command{
	Use: "authcurl",
	Short: "Fetch requests from aws-authenticated API endpoints",
	Run: func (cmd *cobra.Command, args []string) {
		InitializeConfig()
		fmt.Println(viper.Get("Region"))
	}}

var AccessKey, SecretKey, Region string

func InitializeConfig() {
	viper.SetConfigName("authcurl")
	viper.SetEnvPrefix("auth")

	viper.AddConfigPath("$HOME/.authcurl")
	viper.AddConfigPath("$HOME/.config/authcurl")
	viper.AddConfigPath("$HOME/.config")

	viper.SetDefault("Region", "us-east-1")

	viper.BindEnv("region")
	viper.BindEnv("secret_key")
	viper.BindEnv("access_key")

	if authCmdV.PersistentFlags().Lookup("region").Changed {
		viper.Set("Region", Region)
	}
}

func Execute() {
	AuthCurlCmd.Execute()
}

func init() {
	AuthCurlCmd.PersistentFlags().StringVarP(
		&AccessKey, "access-key", "", "", "Your access key to the API")
	AuthCurlCmd.PersistentFlags().StringVarP(
		&SecretKey, "secret-key", "", "", "Your secret key to the API")
	AuthCurlCmd.PersistentFlags().StringVarP(
		&Region, "region", "r", "us-east-1", "The region.  Not terribly useful")
	authCmdV = AuthCurlCmd
}
