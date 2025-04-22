// Write a function that takes an array of integers and returns the sum of all elements.

package main

import "fmt"

func sumArray(arr [5]int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5} //sum 15
	fmt.Println("Sum:", sumArray(arr))
}
