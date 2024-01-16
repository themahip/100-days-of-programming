package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	//initializing the map with make-> helps to allocate the memory for map
	languages := make(map[int]string)
	//adding the value to the key in maps
	languages[1] = "Rust"
	languages[2] = "Go Lang"
	languages[3] = "Java"

	fmt.Println("list of all the languages: ", languages) // printing the maps
	fmt.Println("the first language is: ", languages[01]) // printinf single items fo the maps

	delete(languages, 3)
	fmt.Println("list of all the language is: ", languages)

	// we can also loop the map

	for key, value := range languages {
		fmt.Printf("the key %v language is %v\n", key, value)
	}

	// using comma, ok syntax

	for _, value := range languages {
		fmt.Println("the value is :", value)
	}

}
