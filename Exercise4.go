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

func main() {
	shape := os.Args[1]
	if shape == "circle" {
		r := os.Args[2]
		radius, _ := strconv.ParseFloat(r, 64)
		circle := Circle{radius}
		area := circle.Area()
		fmt.Println("Area:", area)
		perimeter := circle.Perimeter()
		fmt.Println("Perimeter:", perimeter)
	} else {
		w := os.Args[2]
		h := os.Args[3]
		height, _ := strconv.ParseFloat(h, 64)
		breadth, _ := strconv.ParseFloat(w, 64)
		rectangle := Rectangle{height, breadth}
		area := rectangle.Area()
		fmt.Println("Area:", area)
		perimeter := rectangle.Perimeter()
		fmt.Println("Perimeter:", perimeter)
	}
}
