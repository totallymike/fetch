package main

import (
	"github.com/totallymike/aws-authenticated-request-thing/commands"
)

func main() {
	commands.Execute()
	/*
	var AccessKey string
	var SecretKey string

	mainCommand := &cobra.Command{
		Use: "main [url]",
		Short: "Fetch requests from aws-authenticated endpoints",
		Long: "Fetch data from the given url",
		Run: func (cmd *cobra.Command, args []string) {
			req, err := NewSignedRequest("GET")

			resp, err := req.Perform("")
			if err != nil {
				log.Fatal(err)
			}

			robots, err := ioutil.ReadAll(resp.Body)

			resp.Body.Close()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)
		},
	}

	mainCommand.PersistentFlags().StringVarP(
		&AccessKey,
		"access_key",
		"",
		"",
		"your access key to the API")

	mainCommand.PersistentFlags().StringVarP(
		&SecretKey,
		"secret_key",
		"",
		"",
		"your secret key to the API")

	mainCommand.Execute()
	*/
}

