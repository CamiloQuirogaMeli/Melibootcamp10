package main

import (
	"errors"
	"fmt"
)

var monthlySalary float64

func calculateSalary(hoursWorked float64, hourlyWage float64) (float64, error) {

	if hoursWorked < 80 {
		return 0, errors.New("the worker cannot have worked less than 80 hours a month.")
	}

	monthlySalary = hoursWorked * hourlyWage

	if monthlySalary >= 150000 {
		monthlySalary = monthlySalary - monthlySalary*0.1
	}

	return monthlySalary, nil
}

func calculateSemiannualBonus(bestSemesterSalary float64, monthsWorkedSemester float64) (float64, error) {
	if monthsWorkedSemester < 1 {
		return 0, errors.New("the worker cannot have worked less than 1 month in this semester.")
	}

	semiannualBonus := bestSemesterSalary / 12 * monthsWorkedSemester

	return semiannualBonus, nil

}

func handleErrorSalary(monthlySalary float64, err error) {
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("The monthly salary is", monthlySalary)
	}
}

func handleErrorBonus(semiannualBonus float64, err error) {
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("The semiannual bonus is", semiannualBonus)
	}
}
