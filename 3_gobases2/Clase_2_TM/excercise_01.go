package main

import "fmt"

const (
	MINIMUM_TAX = 17.00
	MAXIMUM_MAX = 10.00
)

func main() {

	var salary float64

	fmt.Println("Please, enter your salary: ")
	fmt.Scanln(&salary)

	fmt.Printf("Your tax percetage discount is: [%%%2.f]\n", taxSalaryPercentageDiscount(salary))
}

func taxSalaryPercentageDiscount(salary float64) float64 {

	if salary > 50000.00 && salary < 150000.00 {
		return MINIMUM_TAX
	} else if salary > 150000.00 {
		return MAXIMUM_MAX
	} else {
		return 0
	}

}
