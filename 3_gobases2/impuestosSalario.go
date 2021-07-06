package main

import "fmt"

//Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
//impuesto de un salario. Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del sueldo y si gana más de $150.000 se le descontará además un %10.

func main()  {
	fmt.Println("Se va a calcular el salario total para los empleados 1, 2 y 3...")
	var sal1, sal2, sal3 = 35000.0, 70000.0, 180000.0

	fmt.Println("El empleado 1 tiene un salario neto de $", sal1, "y con los descuentos respectivos, su salario total será $", calImpSal(sal1))
	fmt.Println("El empleado 1 tiene un salario neto de $", sal2, "y con los descuentos respectivos, su salario total será $", calImpSal(sal2))
	fmt.Println("El empleado 1 tiene un salario neto de $", sal3, "y con los descuentos respectivos, su salario total será $", calImpSal(sal3))
}

func calImpSal(sal float64) float64{

	if sal > 150000 {
		return sal * 0.73
	} else if sal > 50000 {
		return sal * 0.83
	} else {
		return sal
	}

}