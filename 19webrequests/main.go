package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web request in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the type of response is: %T", response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("the data is,: \n", string(data))

}
