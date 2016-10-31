package main

import (
	"fmt"
)

var global string = "Hello World from global var"

func main() {
	var x string = "hello world" // var <var_name> type = <value>
	fmt.Println(x)
	x = "second string"
	fmt.Println(x)

	var x2 string
	x2 = "hello world"
	fmt.Println(x2)

	x = x + x2
	fmt.Println(x)

	x += "yoo"
	fmt.Println(x)

	var y string = "xxx"
	fmt.Println(x == y)

	// ---------------------------
	fmt.Println("*****************************")

	yy := "hello world" //Go Complier knows how to infer the type based on the literal value
	fmt.Println(yy)

	var yyy = "hello world2" // is the same as  yyy := "hello world2"
	fmt.Println(yyy)

	xxx := 5
	fmt.Println(xxx)

	dogName := "Max"
	fmt.Println("My dog name is", dogName)

	// ---------------------------
	fmt.Println("*****************************")

	f() // calling a function that access a global var

	// -- constants
	const xConst string = "const hello world" // const <const_name> type = <const_value>
	fmt.Println(xConst)

	// ---------------------------
	fmt.Println("*****************************")
	var ( // this is how to declar muliplte variables. you can use const rather than var
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a + b + c)
	// ---------------------------
	fmt.Println("*****************************")

	fmt.Print("Enter a number")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("the output =", output)
}

func f() {
	fmt.Println(global)
}
