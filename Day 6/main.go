//Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”.

package main

import "fmt"

func fizzBuzz(number int) interface{} {
	switch {
	case (number%3 == 0) && (number%5 == 0):
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	default:
		return number
	}
}

func main() {
	var number int = 1
	for number <= 100 {
		fmt.Println(fizzBuzz(number))
		number += 1
	}
}
