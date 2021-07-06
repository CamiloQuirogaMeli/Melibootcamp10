package main

import (
	"fmt"
	"errors"
)

func main(){
	company := Company{make(map[string]int)}
	company.AddEmployee("Ruperto", 70)

	fmt.Println("* caso salario mensual con horas trabajadas invalidas")
	_, err := company.EmployeeSalary("Ruperto", 40)
	fmt.Println("\t", err)
	
	fmt.Println("* caso salario mensual con horas trabajadas validas")
	salary, _ := company.EmployeeSalary("Ruperto", 100)
	fmt.Println("\t el salario mensual es:", salary)
	
	fmt.Println("* caso medio aguinaldo con horas trabajadas invalidas")
	_, err = company.HalfBonus("Ruperto", 1, 5)
	fmt.Println("\t", err)
	
	fmt.Println("* caso medio aguinaldo con meses trabajadas invalidos")
	_, err = company.HalfBonus("Ruperto", 100, -1)
	fmt.Println("\t", err)
	
	fmt.Println("* caso medio aguinaldo con meses trabajadas invalidos y salario mensual con horas trabajadas invalidas ")
	_, err = company.HalfBonus("Ruperto", 1, -1)
	fmt.Print("\terrores 1 & 2:\n\t",err)
	fmt.Println("\terror 1:",errors.Unwrap(err))
	
	fmt.Println("* caso medio aguinaldo con meses trabajadas validos y con horas trabajadas validas")
	salary, _ = company.HalfBonus("Ruperto", 100, 4)
	fmt.Println("\t", salary)
}