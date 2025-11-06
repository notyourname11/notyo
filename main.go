package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2.5},
	}

	var total float64
	for _, s := range shapes {
		fmt.Println("Площадь:", s.Area())
		total += s.Area()
	}
	fmt.Println("Общая площадь:", total)
}
