/* Write a program that:
Asks the user to enter a number between 0 and 100.
Based on the number, prints out the grade using the following rules:
90-100: Grade A
80-89: Grade B
70-79: Grade C
60-69: Grade D
<60: Grade F
If the number is outside 0–100, print “Invalid input”. */

package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter a score (0-100):")
	fmt.Scanln(&score)

	if score < 0 || score > 100 {
		fmt.Println("Invalid input")
	} else if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else if score >= 60 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Grade F")
	}

}
