package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	greeter()
	adder := adder(3, 4)
	fmt.Println(adder)
	proResult, Message := proAdder(3, 4, 5, 6, 6, 7, 7, 8)
	fmt.Println(proResult)
	fmt.Println(Message)

}

func adder(a int, b int) int {
	total := a + b
	return total

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, i := range values {
		total += i
	}
	return total, "hello i am a pro adder"
}

func greeter() {
	fmt.Println("hello and namaste everyone")
}
