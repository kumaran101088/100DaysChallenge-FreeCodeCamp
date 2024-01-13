// Write a program, which will find all the numbers between 1000 and 3000 (both included) such that each digit of a number is an even number. The numbers obtained should be printed in a comma separated sequence on a single line.

package main

import (
	"fmt"
	"strconv"
)

func evenOrOdd(number string, output *[]int) {
	even := true
	for _, value := range number {
		converted, _ := strconv.Atoi(string(value))
		if converted%2 != 0 {
			even = false
		}
	}

	if even {
		converted, _ := strconv.Atoi(number)
		*output = append(*output, converted)
	}
}

func main() {
	var startEnd int = 1000
	output := make([]int, 0)
	for startEnd <= 3000 {
		evenOrOdd(strconv.Itoa(startEnd), &output)
		startEnd += 1
	}
	fmt.Println(output)
}
