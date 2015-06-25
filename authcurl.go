package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	req, err := NewSignedRequest("GET", "http://localhost:49157/v1/network-hosts")


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

