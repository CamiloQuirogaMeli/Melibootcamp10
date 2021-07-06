package main

import (
	"errors"
	"fmt"
	"os"
)

func getNetSalary(grossSalary float64, workedHs float64) (netSal float64, err error) {
	if workedHs < 80 {
		return 0, errors.New("can't work less than 80hs")
	}

	var tax float64
	if grossSalary > 150000 {
		tax = 10
	}
	if tax != 0 {
		toSub := (grossSalary * tax) / 100
		netSal = grossSalary - toSub
		return netSal, nil
	}
	return grossSalary, nil
}

func bonusSalary(mmWorked int, bestSalary float64) (float64, error) {
	if mmWorked < 0 || bestSalary < 0 {
		return 0, fmt.Errorf("the values are invalid")
	}
	return (bestSalary / 12) * float64(mmWorked), nil
}

func main() {
	var salary float64 = 145000
	var hours float64 = 160
	netSalary, err := getNetSalary(salary, hours)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("Salary to pay: $%.2f\n", netSalary)
	mmWorked := 12
	bestSalary := 150000.0
	bonus, err := bonusSalary(mmWorked, bestSalary)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("Bonus to pay: $%.2f\n", bonus)
}
