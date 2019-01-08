package main

import (
	"fmt"
	"net/url"
)

func main() {
	url := url.URL{
		Host:   "localhost",
		Path:   "/test",
		Scheme: "http",
	}

	fmt.Println(url.RequestURI())
	fmt.Println(url.String())
}
