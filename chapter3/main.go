package main

import (
	"fmt"
)

func main() {
	fmt.Println(`1 + 1 = \n \t`, 1.0+1.0)
	fmt.Println("1 + 1 = \n \t", 1.0+1.0)
	fmt.Println("the length of hello world =", len("Hello World"))
	fmt.Println("hello world"[0])
	fmt.Println("hello " + "world")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
