package main

import (
	"fmt"
	"math"
)

type Shape interface { // structs just need to implement all methods inside the interface (like Rectangle and Circle)
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleAreaByRef(c *Circle) float64 { // this is better since you don't want to copy the var C
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

type Android struct { // here we embded Person into Android so Android is a person and more
	Person
	model string
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {

	//	var c Circle // an instance var of typr Circle

	//	x := new(Circle) // pointer to an object of type circle

	c := Circle{0, 0, 5}

	fmt.Println(c.x, c.y, c.r)

	c.x = 2

	fmt.Println(circleArea(c))

	fmt.Println(circleAreaByRef(&c))

	fmt.Println(c.area())

	r := Rectangle{0, 0, 5, 5}

	fmt.Println(r.area())

	p := Person{name: "Ahmad"}

	p.talk()

	a := new(Android) // is a pointer
	a.name = "Android"
	(*a).name = "x"
	a.talk()

	fmt.Println(totalArea(&c, &r))
}
