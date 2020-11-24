package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chevkforvaccine checks a bool value to return true or false
func checkforvaccine() bool {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	check := (r1.Intn(8))
	if check == 3 {
		return true
	}
	return false

}

func main() {

	var covidexists bool = true
	var cure bool = false
	var infectionrate float64 = 0

	for covidexists {
		infectionrate += 0.003
		// fmt.Println("cases rising")

		if infectionrate >= 0.025 {
			infectionrate = 0.01
			fmt.Println("lock down everything")
			cure = checkforvaccine()
		} else if cure {
			fmt.Println("cure found: covid over")
			covidexists = false

		} else {
			cure = checkforvaccine()
			fmt.Println("doing fuck all")
		}

	}

}
