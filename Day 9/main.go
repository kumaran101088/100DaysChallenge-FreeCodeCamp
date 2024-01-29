//Write  a  program  to  compute  1/2+2/3+3/4+...+n/n+1  with  a  given  n  input  by console (n>0).Example:If the following n is given as input to the program:5 Then, the output of the program should be:3.55

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errorMessage(err error) {
	if err != nil {
		fmt.Println("Error is: ", err)
		os.Exit(1)
	}
}

func calculate(number int) float64 {
	var answer float64 = 0
	for x := 1; x <= number; x++ {
		answer += float64(x) / float64((x + 1))
	}
	return answer
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errorMessage(err)
	number, err := strconv.Atoi(strings.Trim(input, "\n"))
	errorMessage(err)
	fmt.Printf("Output is: %.2f\n", calculate(number))
}
