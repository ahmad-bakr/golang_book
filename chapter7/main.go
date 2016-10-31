package main

import (
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println(average(xs))

	// create a function inside a fucntion
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
