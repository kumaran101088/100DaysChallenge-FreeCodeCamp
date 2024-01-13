// Write a code which accepts a sequence of words as input and prints the words in a sequence after sorting them alphabetically.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	for {

		fmt.Print("Enter a new sequence: ")

		reader := bufio.NewReader(os.Stdin)

		sequence, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error %v", err)
			return
		}

		toSlice := strings.Split(sequence, " ")

		sort.Strings(toSlice)

		sortedSequence := strings.Join(toSlice, " ")

		fmt.Print("Sorted sequence: " + sortedSequence + "\n")

	}

}
