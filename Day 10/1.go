//With   two   given   slices   {1,3,6,78,35,55}   and   {12,24,35,24,88,120,155},   write   a program to make a slice whose elements are intersection of the above given slices.

package main

import "fmt"

func checkIntersection(sliceOne []int, sliceTwo []int) []int {
	answer := make([]int, 0, 0)
	mapper := make(map[int]bool)
	for _, i := range sliceOne {
		for _, j := range sliceTwo {
			if i == j && !mapper[i] {
				answer = append(answer, i)
				mapper[i] = true
			}
		}
	}
	return answer
}

func main() {

	sliceOne := []int{1, 3, 6, 78, 35, 55}
	sliceTwo := []int{12, 24, 35, 24, 88, 120, 155}

	fmt.Println(checkIntersection(sliceOne, sliceTwo))
}
