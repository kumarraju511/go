// sorting float64 {20.5, 18.2, 25.0, 11.11, 55.2, 54.6}

package main

import (
	"fmt"
	"sort"
)

func main() {
	variable := []float64{20.5, 18.2, 25.0, 11.11, 55.2, 54.6}
	fmt.Println("Before sorting float64:", variable)
	sort.Float64s(variable)
	fmt.Println("After sorting float64:", variable) //[11.11 18.2 20.5 25 54.6 55.2]
}
