package main

import "fmt"

func main() {
	fmt.Println("welcome to golang loops")
	// days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	// fmt.Println(days)

	// old ways
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//better ways
	// for i := range days {
	// 	fmt.Println(i)
	// 	fmt.Println(days[i])
	// }

	// we can also do

	// for index, day := range days {
	// 	fmt.Printf("the index is %v and the value is %v\n", index, day)
	// }

	//usin continue inside the loops and we can also use break to break the loop and go out of that scope

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 2 {
			goto here
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("the value is:   ", rogueValue)
		rogueValue++
	}

here:
	fmt.Println("we reached here")

}
