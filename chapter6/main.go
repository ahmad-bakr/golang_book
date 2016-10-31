package main

import (
	"fmt"
)

func main() {
	var x [5]int // x is array of 5 elements

	x[4] = 100
	x[3] = 1
	fmt.Println(x)
	fmt.Println(x[4])

	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	fmt.Println(sum)

	fmt.Println(float64(sum) / float64(len(x)))

	fmt.Println("*******************")
	for i, value := range x {
		fmt.Print(i)
		fmt.Println(value)
	}

	fmt.Println("*******************")
	for _, value := range x { // we tell the complier that we don't need the index i
		fmt.Println(value)
	}

	fmt.Println("*******************")
	y := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(y)

	/// Arrays in GO have fixed length so if you want the length to be changable, you can use slice
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
	fmt.Println(e)

	fmt.Println("*******************")

	//s := make([]int, 5, 5)
	s := []int{0, 1, 2, 3, 4}
	s2 := s
	fmt.Println(cap(s)) // ... 5
	fmt.Println(len(s)) // .. from 0 to 4
	fmt.Println(s)

	s = s[2:4]
	fmt.Println(s)
	fmt.Println(len(s)) // .. 2 and 3
	fmt.Println(cap(s)) // .. from 2 to 4

	s = s[:cap(s)] // .. from 2 to 5
	fmt.Println(s)
	s = s[0:cap(s)]
	s = s2[0:cap(s2)]
	fmt.Println(s)

	fmt.Println("*******************") // copy slices into bigger slices
	t := make([]int, len(s), (cap(s)+1)*2)
	fmt.Println(cap(s))
	fmt.Println(cap(t))
	copy(t, s)
	s = t
	fmt.Println(s)
	fmt.Println(cap(s))

	fmt.Println("*******************") // append data
	p := []byte{2, 3, 5}
	p = append(p, 7, 11, 13)
	// p == []byte{2, 3, 5, 7, 11, 13}

	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}

	fmt.Println("*******************") // maps

	m := make(map[string]int) // key is string and value is int
	m["key"] = 10
	fmt.Println(m)
	fmt.Println(m["key"])
	name, ok := m["key"]
	fmt.Println(name, ok)

	delete(m, "key")
	fmt.Println(m)

	name, ok = m["key"]
	fmt.Println(name, ok)

	if name, ok = m["Un"]; ok {
		fmt.Println(name, ok)
	}

}
