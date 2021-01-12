package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Calculate(s Shape) {
	area := s.Area()
	fmt.Println("Area:", area)
	perimeter := s.Perimeter()
	fmt.Println("Perimeter:", perimeter)
}

func main() {
	shape := os.Args[1]
	if shape == "circle" {
		r := os.Args[2]
		radius, _ := strconv.ParseFloat(r, 64)
		circle := Circle{radius}
		Calculate(circle)
	} else {
		l := os.Args[2]
		b := os.Args[3]
		length, _ := strconv.ParseFloat(l, 64)
		breadth, _ := strconv.ParseFloat(b, 64)
		rectangle := Rectangle{length, breadth}
		Calculate(rectangle)
	}
}
