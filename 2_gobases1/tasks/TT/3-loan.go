package main

import "fmt"

func loan(age int, employ bool, seniority int) (ageR bool, senR bool, empR bool) {
	ageR = false
	senR = false
	empR = false

	/* switch {
	case age > 22:
		fmt.Println("You accomplish with the age requirement:", age, "years old")
		ageR = true
		fallthrough
	case employ == true:
		fmt.Println("You accomplish the employment requirement")
		empR = true
	case seniority > 12:
		fmt.Println("You accomplish the seniority requirement")
		senR = true
	} */

	if age > 22 {
		fmt.Println("You accomplish eith the age requirement:", age, "years old")
		ageR = true
	}

	if employ == true {
		fmt.Println("You accomplish the employment requirement")
		empR = true
	}

	if seniority > 12 {
		fmt.Println("You accomplish the seniority requirement")
		senR = true
	}

	return ageR, senR, empR
}
