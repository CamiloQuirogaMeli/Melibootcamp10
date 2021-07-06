package main

import (
	"errors"
	"fmt"
	"strconv"
)

func applyDiscount(p, d float64) (price float64) {
	toSub := (p * d) / 100
	price = p - toSub
	return price
}

// parseStdIn return the string from stdin to float64
func parseStdIn(in string) (value float64, err error) {
	value, err = strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("There was an error processing the entered value: %v", err))
	}
	return value, nil
}

func main() {
	var in string
	fmt.Print("Enter price($): ")
	fmt.Scanln(&in)
	fullPrice, err := parseStdIn(in)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter discount(%): ")
	fmt.Scanln(&in)
	discount, err := parseStdIn(in)
	if err != nil {
		panic(err)
	}

	price := applyDiscount(fullPrice, discount)
	fmt.Printf("To pay: $%.2f\n", price)
}
