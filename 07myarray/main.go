package main

import "fmt"

func main() {
	fmt.Println("welcone to array of GoLang")

	//fruits list
	var fruitsList [04]string
	fruitsList[0] = "apple"
	fruitsList[1] = "banana"
	fruitsList[3] = "mango"

	fmt.Println("the list of fruits are: ", fruitsList)
	fmt.Println("the number of fruits are: ", len(fruitsList))

	//veg list
	vegList := [3]string{"carrot", "cowliflower", ""}

	fmt.Println(vegList)

}
