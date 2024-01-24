package main

import (
	"fmt"
	"math"
	"strconv"
)

func contains(number int, numbers []int) bool {
	present := false
	for _, value := range numbers {
		if value == number {
			present = true
			break
		}
	}
	if present {
		return true
	}
	return false
}

func ifHappy(number int, numbers []int) bool {
	stringNumber := strconv.Itoa(number)
	var sum int
	for _, letter := range stringNumber {
		numberString, _ := strconv.Atoi(string(letter))
		sum += int(math.Pow(float64(numberString), 2))
	}
	if contains(sum, numbers) {
		if sum == 1 {
			return true
		}
		return false
	}
	numbers = append(numbers, sum)
	return ifHappy(sum, numbers)

}

func main() {

	for x := 1; x <= 100; x++ {
		if ifHappy(x, []int{}) {
			fmt.Printf("%v is a happy number\n", x)
		} else {
			fmt.Printf("%v is not a happy number\n", x)
		}
	}

}
