package main

import "fmt"

func main() {

	var salary float64
	fmt.Println("Ingrese el sueldo: ")
	fmt.Scanf("%f", &salary)
	taxesPerc, taxes := wageTaxes(salary)
	fmt.Printf("El porcentaje es %.2f y el descuento a aplicar es %.2f", taxesPerc, taxes)
}

func wageTaxes(salary float64) (float64, float64) {
	if salary >= 150000 {
		return 27, salary * 27 / 100
	} else if salary > 50000 && salary < 150000 {
		return 17, salary * 17 / 100
	} else {
		return 0, salary * 0 / 100
	}

}
