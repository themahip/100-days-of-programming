package main

import (
	"bufio"
	"fmt"
	"os"
)

type keyValue struct {
	key   int
	value string
}

func main() {
	var keyValueSlice []keyValue
	var option int

	fmt.Printf("***** Welcome to toDo List Maker ******** \n\n")
	fmt.Println("what do you want to do? read or right?")
	fmt.Printf("1.Read\n2.write\ntype 1 for read and 2 for write\n")
	_, err := fmt.Scan(&option)
	if err != nil {
		panic(err)
	}
	for option != 1 && option != 2 {
		fmt.Println("please re enter the value either 1 or 2:")
		fmt.Scan(&option)
	}

	switch option {
	case 1:
		read()

	case 2:
		write(keyValueSlice)

	}
}

func write(s []keyValue) {
	var newItem keyValue //type of newItem
	fmt.Println("what do you want to add? write as '1 'hello world''")
	_, err := fmt.Scan(&newItem.key)
	inputReader := bufio.NewReader(os.Stdin)
	newItem.value, _ = inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s = append(s, newItem)
	content := []byte(fmt.Sprintf("%d,%s", newItem.key, newItem.value))
	err = os.WriteFile("todo.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}

func read() {
	content, err := os.ReadFile("todo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
