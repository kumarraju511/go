/* Write a Go program that:
Asks the user to enter their age.
Uses if-else to print:
"You are a minor" if age < 18
"You are an adult" if age is between 18 and 64
"You are a senior" if age ≥ 65
"Invalid age" if age is negative*/

package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 0 {
		fmt.Println("Invalid age")
	} else if age < 18 {
		fmt.Println("You are a minor")
	} else if age < 65 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior")
	}
}
