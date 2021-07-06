package main

import "fmt"

func main() {
	result := salaryAfterTaxes(60000)
	fmt.Println(result)
}

func salaryAfterTaxes(salary float64) float64 {
	var tax float64 = 0.0
	if salary > 50000 {
		tax = tax + 17
	}
	if salary > 150000 {
		tax = tax + 10
	}
	result := salary * (1 - (tax / 100))
	return result
}
