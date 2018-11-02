package main

import (
	"fmt"
)

func main() {

	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.1415139475

	//Declaring multiple variables
	x, y := 14, 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(y)

	xx, yy := 5, 6

	fmt.Println("x + yy = ", xx+yy)
	fmt.Println("x - yy = ", xx-yy)
	fmt.Println("x * yy = ", xx*yy)
	fmt.Println("x / yy = ", xx/yy)
	fmt.Println("x mod yy = ", xx%yy)

	isbool := true
	hate := false

	fmt.Println(isbool && hate)
	fmt.Println(isbool || hate)
	fmt.Println(!isbool)

}
