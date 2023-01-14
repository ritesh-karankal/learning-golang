package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The response is of type: %v\n", response)

	defer response.Body.Close()

	datatypes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	content := string(datatypes)
	fmt.Println(content)
}
