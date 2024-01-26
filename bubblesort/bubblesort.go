package main

import "fmt"

func bubbleSort(arr []int) {

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}
		i++
	}

}

func main() {
	arr := []int{1, 78, 32, 8, 5, 6, 53, 22, 9, 0}
	bubbleSort(arr)
	fmt.Println(arr)

}
