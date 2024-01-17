package main

import "fmt"

func main() {
	fmt.Println("if else in golang ")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "number is less than 10"
	} else if loginCount == 10 {
		result = "number is equal"
	} else {
		result = "number is greater"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	// initializing the variable inside the if else statement it self
	if num := 3; num < 10 {
		fmt.Println("num is smaller than 10 ")
	}
}
