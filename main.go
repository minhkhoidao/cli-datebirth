package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for the date of birth
	fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Clean up input
	input = strings.TrimSpace(input)

	// Parse the input date
	dob, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	// Calculate the age
	currentTime := time.Now()
	age := currentTime.Year() - dob.Year()
	if currentTime.Month() < dob.Month() || (currentTime.Month() == dob.Month() && currentTime.Day() < dob.Day()) {
		age--
	}

	// Print the age
	fmt.Printf("You are %d years old.\n", age)
}
