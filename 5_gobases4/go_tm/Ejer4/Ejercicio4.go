/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
a) Salario mensual de un trabajador según la cantidad de horas trabajadas. En caso de
que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10%
en concepto de impuesto. La función que se ocupe de realizar este cálculo deberá
retornar más de un valor, incluyendo un error en caso de que la cantidad de horas
mensuales ingresadas sea menor a 80 o un número negativo. El error deberá indicar
“error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
b) Calcular el medio aguinaldo correspondiente al trabajador (fórmula de cálculo de
aguinaldo: [mejor salario del semestre] dividido 12 y multiplicar el [resultado
obtenido] por la [cantidad de meses trabajados en el semestre]). La función que
realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que
se ingrese un número negativo.
2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando
“errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los
retornos de error en tu función “main()”.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type CustomError struct {
}

func (e CustomError) Error() string {
	return "error: el trabajador no puede haber trabajado menos de 80 hs mensuales"
}

func calcularSalario(hours int) (float64, float64, error) {
	if math.Signbit(float64(hours)) {
		return 0, 0, errors.New("No se admiten numeros negativos")
	}
	pagoHora := 1875

	total := pagoHora * hours
	var day float64
	day = float64(hours) / 24
	if total > 15000 {
		impuesto := (total * 10) / 100
		total = total - impuesto
		return float64(total), day, nil
	} else {
		return float64(total), day, nil
	}

}

func calcularAguinaldo(salary float64, months int) (float64, error) {
	if math.Signbit(float64(salary)) || math.Signbit(float64(months)) {
		return 0, errors.New("No se admiten numeros negativos")
	}

	aguinaldo := (salary * 12) * float64(months)

	return aguinaldo, nil
}

func main() {
	const OPTIONS = `
	1. Calcular salario
	2. Calcular medio aguinaldo
	3. Salir
	`
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenido a su programa de nomina")
	fmt.Println()
	fmt.Println("Selecione la opcion que desee ejecutar")
	fmt.Println(OPTIONS)
	var option int
	_, errOption := fmt.Scanf("%d", &option)

	if errOption != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}

	switch option {
	case 1:
		fmt.Println("¿Cuantas horas trabajo en el mes?")
		scanner.Scan()
		hours, err := strconv.Atoi(scanner.Text())
		if err != nil || (hours < 80 && hours > 0) {
			errorHours := CustomError{}
			errorData := fmt.Errorf("No se ingreso un dato valido o %w", errorHours)

			if hours < 80 {
				fmt.Println(errors.Unwrap(errorData))
			} else {
				fmt.Println(errorData)
			}
		} else {
			total, days, err := calcularSalario(hours)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Su salario es de:", total, "Los dias que trabajo fueron:", days)
			}

		}

	case 2:
		fmt.Println("Para calcular el medio aguinaldo, es necesario que digite los siguientes datos:")
		fmt.Println("Mejor salario del semestre")
		scanner.Scan()
		salary, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("La data ingresada no es valida")
		} else {
			fmt.Println("Cuantos meses trabajo en la compañia")
			scanner.Scan()
			months, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("La data ingresada no es valida")
			} else {
				aguinaldo, err := calcularAguinaldo(salary, months)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Su aguinaldo es: %.2f", aguinaldo)
				}
			}

		}

	default:
		fmt.Println("Gracias por utilizar el programa")
		os.Exit(0)
	}
}
