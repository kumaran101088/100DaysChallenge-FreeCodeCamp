//Write a program to print the slice after removing the value 24 in {12,24,35,24,88,120,155}

package main

import "fmt"

func afterRemoving(sliceOne []int, number int) []int {
	newSlice := make([]int, 0, 0)
	for _, value := range sliceOne {
		if number != value {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func main() {

	sliceOne := []int{12, 24, 35, 24, 88, 120, 155}

	fmt.Println(afterRemoving(sliceOne, 24))
}
