package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 1
	for x <= 10 {
		fmt.Println(x)
		x++
	}

	for a := 1; a < 5; a++ {
		fmt.Println(a)
		for b := 1; b < 5; b++ {
			fmt.Println("*")
		}
	}
}
