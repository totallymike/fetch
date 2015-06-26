package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
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
}

