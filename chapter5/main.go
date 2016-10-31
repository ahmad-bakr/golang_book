package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("*************")

	for x := 1; x < 10; x++ {
		if x%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	fmt.Println("*************")

	y := 3

	switch y {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("undefined")
	}

}
