package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Customer struct {
	Name    string
	Email   string
	Gender  string
	Age     int
	Tickets int
}

func ageValidator(age int) bool {
	if age >= 18 {
		return true
	}
	fmt.Println("You need to be 18 or above to view this show")
	return false
}

func emailValidator(email string) bool {
	valid := false
	if strings.Contains(email, "@") {
		if strings.Contains(strings.Split(email, "@")[1], ".") {
			if strings.Count(strings.Split(email, "@")[1], ".") == 1 {
				valid = true
				return valid
			}
		}
	}
	fmt.Println("Please enter a valid email address")
	return valid
}

func genderValidator(gender string) bool {
	genders := []string{"m", "f", "t"}
	for _, value := range genders {
		if string(value) == strings.ToLower(gender) {
			return true
		}
	}
	fmt.Println("Please enter a valid gender")
	return false
}

func ticketValidator(tickets int, ttickets int) bool {
	if tickets <= ttickets {
		return true
	}
	fmt.Println("Requested amount of tickets not available.")
	return false
}

func masterValidator(age int, gender string, email string, tickets int, ttickets int) bool {
	if ageValidator(age) && genderValidator(gender) && emailValidator(email) && ticketValidator(tickets, ttickets) {
		fmt.Println("Valid")
		return true
	} else {
		fmt.Println("Please try again")
		return false
	}
}

func addToStruct(name string, email string, gender string, age int, tickets int, data *[]Customer) {
	var customer Customer
	customer.Name = name
	customer.Email = email
	customer.Gender = gender
	customer.Age = age
	customer.Tickets = tickets
	*data = append(*data, customer)
}

func gatherInfo() (string, string, string, int, int) {
	var email string
	var gender string
	var age int
	var tickets int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter your gender (m/f/t): ")
	fmt.Scan(&gender)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Print("No. of tickets needed: ")
	fmt.Scan(&tickets)
	return strings.ToTitle(strings.Trim(name, " ")), strings.Trim(email, " "), strings.ToTitle(strings.Trim(gender, " ")), age, tickets
}

func main() {
	var totalTickets int = 50
	var data []Customer
	for {
		fmt.Println("Available tickets: ", totalTickets)
		if totalTickets > 0 {
			name, email, gender, age, tickets := gatherInfo()
			if masterValidator(age, gender, email, tickets, totalTickets) {
				addToStruct(name, email, gender, age, tickets, &data)
				totalTickets -= tickets
				fmt.Println(data)
			}
		} else {
			fmt.Println("Tickets unavailable for the event")
			break
		}
	}
}
