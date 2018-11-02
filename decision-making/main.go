package main

import (
	"fmt"
)

func main() {
	age := 20

	if age > 16 {
		fmt.Println("Yes, you can vote!")
	} else {
		fmt.Println("No, you can't vote")
	}

	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 18:
		fmt.Println("Don't run after girls")
	case 20:
		fmt.Println("Get yourself a job!")
	default:
		fmt.Println("Are you even alive?")
	}
}
