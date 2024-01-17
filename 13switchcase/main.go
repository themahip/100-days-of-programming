package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and Case in Go lang")
	// var myTime = time.Date(2023, time.December, 31, 12, 0, 0, 0, time.UTC)
	// rand.Seed(myTime.UnixNano())
	diceNumber := rand.Intn(7)
	fmt.Println(diceNumber)

	// once the once case is hit it wont go through other cases but if we want it we can mention fall through
	switch diceNumber {
	case 1:
		fmt.Println("you hit 1")
	case 2:
		fmt.Println("you hit 2")
	case 3:
		fmt.Println("you hit 3")
		fallthrough
	case 4:
		fmt.Println("you hit 4")
	case 5:
		fmt.Println("you hit 5")
	case 6:
		fmt.Println("you hit 6 you can roll again")
	default:
		fmt.Println("what was that!!")
	}
}
