package main

import (
	"fmt"
)

func main() {
	// StudentAge := make(map[string]int)

	// StudentAge["Eduardo"] = 21
	// StudentAge["Ana"] = 30
	// StudentAge["Barbara"] = 20
	// StudentAge["Maria"] = 18

	// fmt.Println(StudentAge)
	// fmt.Println(StudentAge["Eduardo"])
	// fmt.Println(StudentAge["Ana"])
	// fmt.Println(len(StudentAge))

	superHero := map[string]map[string]string{
		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},

		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}

	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}

}
