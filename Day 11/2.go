//Write a program to print the slice after removing the numbers divisble by 5 and 7 in {12,24,35,70,88,120,155}

package main

import "fmt"

func checkDivisble(numbers []int) []int {
	divisibleNumbers := make([]int, 0, 0)
	for _, number := range numbers {
		if (number%5 == 0) && (number%7 == 0) {
			divisibleNumbers = append(divisibleNumbers, number)
		}
	}
	return divisibleNumbers
}

func main() {
	numbers := []int{12, 24, 35, 70, 88, 120, 155}
	fmt.Println(checkDivisble(numbers))
}
