package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	err := makeHTTPRequest()
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}

	err = makeHTTPRequest_fp()
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
}

func makeHTTPRequest() error {
	// ok: missing-defer-http
	resp, err := http.Get("http://example.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response from http://example.com:")
	fmt.Println(string(body))

	return nil
}

func makeHTTPRequest_fp() error {
	// ruleid: missing-defer-http
	resp, err := http.Get("http://example.com")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response from http://example.com:")
	fmt.Println(string(body))

	return nil
}
