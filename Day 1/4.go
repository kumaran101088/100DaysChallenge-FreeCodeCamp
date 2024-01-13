// Write a program that accepts a sentence and calculate the number of letters and digits.

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func countLetters(sentence string) int {
	output := 0
	for _, letter := range sentence {
		if unicode.IsLetter(letter) {
			output += 1
		}
	}
	return output
}

func countDigits(sentence string) int {
	output := 0
	for _, letter := range sentence {
		if unicode.IsDigit(letter) {
			output += 1
		}
	}
	return output
}

func main() {
	fmt.Print("Enter a sentence: ")

	reader := bufio.NewReader(os.Stdin)

	sentence, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	letters := countLetters(sentence)

	digits := countDigits(sentence)

	fmt.Printf("%v contains %v letters and %v digits", sentence, letters, digits)
}
