// Write a program which accepts a string from console and print the characters that have even indexes.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter your input: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error is ", err)
	}
	input = strings.Trim(input, " ")
	for index, value := range input {
		if index%2 == 0 {
			fmt.Printf(string(value))
		}
	}
	fmt.Println("\n")
}
