package main

import "fmt"

func selectionSort(toSort []int) []int {
	for index, value := range toSort {
		min := value
		swapIndex := index
		for j := index + 1; j <= len(toSort)-1; j++ {
			if value > toSort[j] && toSort[j] < min {
				min = toSort[j]
				swapIndex = j
			}
		}
		toSort[index], toSort[swapIndex] = toSort[swapIndex], toSort[index]
	}
	return toSort
}

func bubbleSort(toSort []int) []int {
	for i, _ := range toSort {
		for j := i + 1; j <= len(toSort)-1; j++ {
			if toSort[i] > toSort[j] {
				toSort[i], toSort[j] = toSort[j], toSort[i]
			}
		}
	}
	return toSort
}

func main() {
	toSort := []int{7, 2, 8, 1, 10, 5, 4, 3, 9, 0, 6}
	fmt.Println(selectionSort(toSort))
	fmt.Println(bubbleSort(toSort))
}
