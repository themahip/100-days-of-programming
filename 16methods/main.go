package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	fmt.Println("welcome to methods in golang")
	Mahip := User{"Mahip Adhikari", 21, "mahip.adhikari4@gmail.com", true}
	fmt.Println("About the first user; ", Mahip)
	fmt.Printf("the values are: %+v\n", Mahip)
	Mahip.getStatus()
	Mahip.newMail()

}

// now lets define the fuction

func (u User) getStatus() {
	fmt.Printf("%v active status is %v\n", u.Name, u.Status)
}

func (u User) newMail() {
	u.Email = "mahip.adhikari3@gmail.com"
}
