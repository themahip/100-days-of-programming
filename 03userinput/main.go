package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(welcome)
	fmt.Println("enter the rating for our pizza: ")

	// comma okay syntax || err ok

	input, err := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Println("the error is", err)

}
