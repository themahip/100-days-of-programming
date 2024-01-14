package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza")
	rate := bufio.NewReader(os.Stdin)
	rating, _ := rate.ReadString('\n')
	fmt.Println("the rating is:", rating)
	fmt.Printf("the type of rating is,%T \n", rating)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 32)
	if err != nil {
		println(err)
	} else {
		println("the new rating is", newRating+1)
	}

}
