package main

import (
	"fmt"
	"errors"
)

type Company struct {
	employees map[string]int
}

func (o *Company) AddEmployee(anEmployeeName string, hourlyWage int) {

	o.employees[anEmployeeName] = hourlyWage
}

func (o *Company) EmployeeSalary(anEmployeeName string, hoursWorkedForMonth int) (int, error) {
	
	if hoursWorkedForMonth < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	monthSalary := o.employees[anEmployeeName] * hoursWorkedForMonth
	if monthSalary >= 150000 {
		monthSalary -= int(0.1 * float64(monthSalary))
	}

	return monthSalary, nil
}
 
func (o *Company) HalfBonus(anEmployeeName string, hoursWorkedPerMonth int, monthsWorked int) (int, error) {
	
	highestSemesterSalary, err := o.EmployeeSalary(anEmployeeName, hoursWorkedPerMonth)

	if monthsWorked < 0 {
		if err != nil{
			return 0, fmt.Errorf("error: la cantidad de meses trabajados no puede ser negativa - \n\t%w", err)
		} else {
			return 0, errors.New("error: la cantidad de meses trabajados no puede ser negativa")
		}
	}
	if err != nil {
		return 0, err
	}

	monthSalary := (highestSemesterSalary / 12) * monthsWorked

	return monthSalary, nil
}