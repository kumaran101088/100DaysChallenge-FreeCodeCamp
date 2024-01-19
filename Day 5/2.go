// Write a program which accepts a string from console and print it in reverse order.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your input: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error is ", err)
	}
	var newString string
	for i := len(input) - 1; i >= 0; i-- {
		newString += string(input[i])
	}
	fmt.Println(newString)
}
