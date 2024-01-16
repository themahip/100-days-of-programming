package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Here we are learning structures in go lang")

	employee := User{"Mahip", "abc@gmail.com", 21, true}
	fmt.Println(employee)

	//using +v syntax
	fmt.Printf("the details are: %+v\n", employee)

}
