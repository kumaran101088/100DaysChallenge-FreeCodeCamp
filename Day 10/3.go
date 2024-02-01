//With a given slice {12,24,35,24,88,120,155,88,120,155}, write a program to print this slice after removing all duplicate values with original order reserved

package main

import "fmt"

func makeReverse(sliceOne []int) []int {
	reverseSlice := make([]int, 0, 0)
	for i := 0; i < len(sliceOne); i++ {
		reverseSlice = append(reverseSlice, sliceOne[(len(sliceOne)-1)-i])
	}
	return reverseSlice
}

func checkDuplicates(reversedSlice []int) []int {
	mapper := make(map[int]bool)
	answer := make([]int, 0, 0)
	for _, value := range reversedSlice {
		if !mapper[value] {
			answer = append(answer, value)
			mapper[value] = true
		}
	}
	return answer

}

func afterChanges(sliceOne []int) []int {
	reversedSlice := makeReverse(sliceOne)
	return checkDuplicates(reversedSlice)
}

func main() {

	sliceOne := []int{12, 24, 35, 24, 88, 120, 155, 88, 120, 155}
	fmt.Println(afterChanges(sliceOne))

}
