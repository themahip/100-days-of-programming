package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this is the need to go in the file"
	myFile, err := os.Create("./mygofile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(myFile, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("the length is : ", length)
	readFile("./mygofile.txt")
	defer myFile.Close()

}

func readFile(myFile string) {
	file, err := ioutil.ReadFile(myFile)

	if err != nil {
		panic(err)
	}
	fmt.Println("the data in file is: ", string(file))

}
