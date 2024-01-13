// Write a program which will find factors of given number and find whether the factor is even or odd.

package main

import "fmt"

func findEvenOdd(number int) {
	if number%2 == 0 {
		fmt.Printf("%v is even\n", number)
	} else {
		fmt.Printf("%v is odd\n", number)
	}
}

func printFactors(number int) {
	actualNumber := number
	for number > 0 {
		if number%2 == 0 {
			fmt.Printf("%v is a factor of %v\n", number, actualNumber)
			findEvenOdd(number)
		} else if number == 1 {
			fmt.Printf("1 is a factor of %v\n", actualNumber)
			fmt.Printf("%v is odd\n", number)
		}
		number -= 1
	}
}

func main() {

	fmt.Print("Enter a number: ")

	var number int

	fmt.Scan(&number)

	printFactors(number)

}
