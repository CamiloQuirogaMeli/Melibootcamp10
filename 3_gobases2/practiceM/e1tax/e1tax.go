package main

import "fmt"

func main() {
	var salary float64
	fmt.Printf("Ingrese el salario: ")
	fmt.Scanf("%f\n", &salary)

	tax := taxtCalc(salary)
	fmt.Println("El inpuesto a cobrar es: ", tax)

}

func taxtCalc(salary float64) string {
	var tax string
	switch {
	case salary > 150000:
		{
			tax = "10%"
		}
	case salary > 50000:
		{
			tax = "17%"
		}
	default:
		{
			tax = "0%"
		}
	}
	return tax
}
