// Write a loop to print the multiplication table (1 to 10) for a given number.

package main

import "fmt"

func main() {
	num := 5
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
