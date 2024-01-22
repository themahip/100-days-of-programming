package main

import "fmt"

func main() {
	fmt.Println("welcome to defer in golang")
	defer fmt.Println("world")

	// here all the defer statment are put on stack and when called they will execute by poping before all the statement of the function or method is executed
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	myDefer()
	fmt.Println("hello")

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
