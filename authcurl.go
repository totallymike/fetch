package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"flag"
)

func main() {

	access_key := flag.String("access_key", "", "your access key")
	secret_key := flag.String("secret_key", "", "your secret key")

	flag.Parse()

	os.Setenv("PWNIE_ACCESS_KEY", *access_key)
	os.Setenv("PWNIE_SECRET_KEY", *secret_key)


	url := flag.Arg(0)

	req, err := NewSignedRequest("GET", url)


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
}

