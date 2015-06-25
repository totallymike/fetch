package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		"http://localhost:49157/v1/network-hosts",
		nil,
	)

	req.Header.Add("Content-Type", "application/vnd.api+json")
	fmt.Printf("%s\n", req.URL.Path)
	fmt.Printf("%v\n", req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", resp.Status)
	robots, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

