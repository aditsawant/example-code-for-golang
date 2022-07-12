package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Getting Content-Type value from the HTTP response")

	url := "http://www.linkedin.com"

	ctype, err := getCType(url)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(ctype)
	}
}

func getCType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find content-type in the header")
	}

	return ctype, nil
}
