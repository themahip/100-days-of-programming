package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices in GO lang")
	fruitsList := []string{"apple", "banana", "orrange", "mango"}
	fmt.Printf("the type of fruit list is %T \n", fruitsList)

	var ptr *[]string = &fruitsList
	*ptr = append(*ptr, "peach", "avocado")
	fmt.Println(fruitsList)

	*ptr = append(fruitsList[0:3])
	fmt.Println(fruitsList)

	highScores := make([]string, 4)
	highScores[0] = "hello"
	highScores[1] = "welcome"
	highScores[2] = "to"
	highScores[3] = "Go"
	//gives an error
	//highScores[4]="lang"

	// but this append method wont
	highScores = append(highScores, "lang")
	fmt.Println(highScores)

	// now for sorting

	sort.Strings(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.StringsAreSorted(highScores))
}
