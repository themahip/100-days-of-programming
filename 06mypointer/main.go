package main

import "fmt"

func main() {
	fmt.Println("welcome to the pointer in GOLang")

	//initializing pointer in GO
	// var ptr *int
	// fmt.Println("the value of pointer is: ", ptr)

	myValue := 22
	ptr := &myValue

	fmt.Println("the address is ", ptr)
	fmt.Println("the value stored is ", *ptr)

	*ptr = *ptr + 3

	fmt.Println("the new value is ", myValue)
	fmt.Println("the address of myValue is: ", &myValue)

}
