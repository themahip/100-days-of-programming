package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=learninggo123"

func main() {
	fmt.Println("Urls in golang")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Printf("the type of query param is %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for index, value := range qparams {
		fmt.Printf(" %v : %v", index, value)
	}

}
