package main

import (
	"errors"
	"fmt"
	"os"
)

type client struct {
	age       int
	employee  bool
	salary    float64
	seniority float64
}

// potentialRisk evaluates if the client can apply for a loan
func potentialRisk(c client) (isElegible bool, err error) {
	isElegible = true
	if c.age <= 22 {
		err = errors.New("The age of the client is less or equal than 22")
		return false, err
	}
	if !c.employee {
		err = errors.New("The client is not an employee")
		return false, err
	}
	if c.seniority <= 1 {
		err = errors.New("The client does not meet the seniority requirement")
		return false, err
	}
	return isElegible, nil
}

func main() {
	c := client{
		age:       23,
		employee:  true,
		salary:    110000,
		seniority: 10,
	}
	isElegible, err := potentialRisk(c)
	if !isElegible && err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("The client is elegible to take a loan")

	if c.salary < 100000 {
		fmt.Println("Interest will be applied")
		os.Exit(0)
	}
	fmt.Println("Interest will not be applied")
}
