package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	hours int
	hourValue int
	bestMonthlySalary int
	monthsWorked int
}

func (e Employee) calculateSalary() (salary int, minHoursError error) {
	if e.hours < 80 {
		minHoursError = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		salary = 0
	}else {
		netSalary := e.hourValue * e.hours
		err := checkSalary(netSalary)
		if err != nil{
			salary = netSalary
		}
		salary = netSalary - calculateTax(netSalary)
	}
	return
}

func calculateTax (salary int) int {
	return int(float64(salary) * 0.10)
}

func checkSalary (salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: [%d]", salary)
	}
	return nil
}

func (e Employee) calculateBonus() (bonus int,err error) {
	if e.monthsWorked < 0 || e.bestMonthlySalary < 0 {
		err = errors.New("error: los datos negativos no son validos")
	}
	bonus = int((float64(e.bestMonthlySalary) / 12.0) * float64(e.monthsWorked))
	return
}

func main(){
	employee := Employee{hours: 79,hourValue: 100,bestMonthlySalary: -1,monthsWorked: -1}

	salary, errSalary := employee.calculateSalary()
	bonus, errBonus := employee.calculateBonus()
	if errSalary != nil {
		fmt.Println(errSalary)
	} else {
		fmt.Println("The salary of employee is: ", salary)
	}
	if errBonus != nil {
		fmt.Println(errBonus)
	} else {
		fmt.Println("The bonus of employee is: ", bonus)
	}

}
