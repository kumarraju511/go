// Write a function that reverses the elements of an array in place.

package main

import "fmt"

func reverseArray(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	reverseArray(&arr)
	fmt.Println("Reversed:", arr)
}
