package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghj343455"

func main() {

	fmt.Println(myURL)

	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qparams := result.Query()
	fmt.Printf("The type of query path are: %v\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("The param is:", val)
	}

	partsofURL := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=ritesh",
	}

	anotherURL := partsofURL.String()
	fmt.Println(anotherURL)
}
