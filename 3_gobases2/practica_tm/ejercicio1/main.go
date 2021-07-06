package main

import "fmt"

//
// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
// depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
// impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
// sueldo y si gana más de $150.000 se le descontará además un %10.
//

func impuestoSalario(salario float64) (float64, int) {
	switch {
	case salario > 50000 && salario < 150000:
		return (salario * 0.17), 17
	case salario > 150000:
		return (salario * 0.1), 10
	}
	return 0, 0
}

func main() {

	fmt.Println("===== Impuestos de salario =====")

	// Empleado - Salario
	var empleados = map[string]float64{"Benjamin": 20000, "Nahuel": 55000, "Brenda": 160000, "Darío": 60000, "Pedro": 10000}

	for nombre, sueldo := range empleados {
		impuesto, porcentaje := impuestoSalario(sueldo)
		totalSueldo := sueldo - impuesto
		fmt.Println("Empleado(a):", nombre)
		fmt.Println("Sueldo: $", sueldo)
		fmt.Printf("Porcentaje de impuesto: %d%%\n", porcentaje)
		fmt.Println("Valor del impuesto: $", impuesto)
		fmt.Println("Total: $", totalSueldo)
		fmt.Println("================================")
	}

}
