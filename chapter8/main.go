package main

import (
	"fmt"
)

func zero(xPtr *int) { // xPtr is a pointer to an integer value
	*xPtr = 0 // derefrence xPtr
}

func one(xPtr *int) {
	*xPtre = 1
}

func main() {
	x := 5
	zero(&x) // & indicates the address of the element x
	fmt.Println(x)
	//*****************

	xPtr := new(int)   // the new keyword return a pointer to an integer
	one(xPtr)          // pass the pointer to a function that accetps a pointer
	fmt.Println(*xPtr) // derefernce the pointer

}
