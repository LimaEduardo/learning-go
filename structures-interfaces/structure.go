package main

import (
	"fmt"
)

func main() {

	// rect1 = Rectangle{height: 10, width: 5}
	rect1 := Rectangle{10, 5}

	fmt.Println(rect1.height)
	fmt.Println(rect1.width)

	fmt.Println("Area: ", rect1.getArea())

}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) getArea() float64 {
	return rect.height * rect.width
}
