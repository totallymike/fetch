package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/totallymike/fetch/request"
	"fmt"
	"strings"
	"log"
	"io/ioutil"
)

var authCmdV *cobra.Command
var AuthCurlCmd = &cobra.Command{
	Use: "authcurl [url]",
	Short: "Fetch requests from aws-authenticated API endpoints",
	Run: func (cmd *cobra.Command, args []string) {
		InitializeConfig()
		request, _ := request.NewSignedRequest("GET", strings.Join(args, ""))
		response, err := request.Perform("")
		defer response.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)

		fmt.Println(viper.Get("Region"))
	}}

var AccessKey, SecretKey, Region, Service string

func InitializeConfig() {
	viper.SetConfigName("authcurl")
	viper.SetEnvPrefix("auth")

	viper.AddConfigPath("$HOME/.authcurl")
	viper.AddConfigPath("$HOME/.config/authcurl")
	viper.AddConfigPath("$HOME/.config")

	viper.SetDefault("Region", "us-east-1")
	viper.SetDefault("allow_insecure_ssl", false)
	viper.SetDefault("service", "")

	viper.BindEnv("region")
	viper.BindEnv("secret_key")
	viper.BindEnv("access_key")

	if authCmdV.PersistentFlags().Lookup("region").Changed {
		viper.Set("Region", Region)
	}

	if authCmdV.PersistentFlags().Lookup("access-key").Changed {
		viper.Set("access_key", AccessKey)
	}

	if authCmdV.PersistentFlags().Lookup("secret-key").Changed {
		viper.Set("secret_key", SecretKey)
	}

	if authCmdV.PersistentFlags().Lookup("service").Changed {
		viper.Set("service", Service)
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
	AuthCurlCmd.PersistentFlags().StringVarP(
		&Service, "service", "s", "", "The service name of the endpoint.  Most of the time you won't worry about it")
	authCmdV = AuthCurlCmd
}
