/* Write a Go program that:
Asks the user to enter a number.
Checks if the number is even or odd using if-else.
Prints "Even Number" or "Odd Number" accordingly.

Bonus: Add logic to print "Zero" if the number is 0. */

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num == 0 {
		fmt.Print("Zero")
	} else if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
