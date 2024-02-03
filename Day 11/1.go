//Write a program which count and print the numbers of each character in a string input by console. Example: If the following string is given as input to the program:abcdefgabc Then, the output of the program should be: a2,c2,b2,e1,d1,g1,f1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printChars(charsDict map[string]int) {
	for key, value := range charsDict {
		fmt.Printf("%v,%v\n", key, value)
	}
}

func countChars(content string) {
	charsDict := make(map[string]int)
	for _, value := range content {
		_, found := charsDict[string(value)]
		if found {
			charsDict[string(value)] += 1
		} else {
			charsDict[string(value)] = 1
		}
	}
	printChars(charsDict)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading input", err)
	} else {
		content = strings.Trim(content, "\n")
		countChars(content)
	}
}
