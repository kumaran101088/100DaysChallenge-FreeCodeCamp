// A website requires a user to input username and password to register. Write a program to check the validity of password given by user. Following are the criteria for checking password:
// 1. At least 1 letter between [a-z]
// 2. At least 1 number between [0-9]
// 3. At least 1 letter between [A-Z]
// 4. At least 1 character from [$#@]
// 5. Minimum length of transaction password: 6
// 6. Maximum length of transaction password: 12

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func spaceChecker(password string) bool {
	if strings.Contains(password, " ") {
		return false
	}
	return true
}

func lengthChecker(password string) bool {
	if len(password) >= 6 && len(password) <= 12 {
		return true
	}
	return false
}

func otherConditionChecker(password string) bool {
	lowerChecker := false
	upperChecker := false
	digitChecker := false
	splCharChecker := false
	for _, letter := range password {
		if !lowerChecker || !upperChecker || !digitChecker || !splCharChecker {
			if unicode.IsLower(letter) {
				lowerChecker = true
			} else if unicode.IsUpper(letter) {
				upperChecker = true
			} else if unicode.IsDigit(letter) {
				digitChecker = true
			} else if (string(letter) == "$") || (string(letter) == "#") || (string(letter) == "@") {
				splCharChecker = true
			}
		}
	}
	if lowerChecker && upperChecker && digitChecker && splCharChecker {
		return true
	}
	return false
}

func main() {

	for {
		fmt.Println("\n######### ######### ######### #########")
		fmt.Print("\nEnter password to validate: ")

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error message is ", err)
			return
		}

		if !lengthChecker(input) {
			fmt.Println("\nYour password should be within 6 to 12 characters in length")
			continue
		}

		if !spaceChecker(input) {
			fmt.Println("\nYour password should not have spaces")
			continue
		}

		if otherConditionChecker(input) {
			formattedString := fmt.Sprintf("\nYour password %s is VALID\n", input)
			fmt.Printf(formattedString)
		} else {
			fmt.Printf("\nYour password '%s' is INVALID\n", input)
			fmt.Println("\nMake sure your password has\n1. At least 1 letter between [a-z]\n2. At least 1 number between [0-9]\n3. At least 1 letter between [A-Z]\n4. At least 1 character from [$#@]")
		}
	}
}
