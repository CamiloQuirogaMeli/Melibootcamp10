package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	salarioMensual, err := calcularSalarioMensual(600)
	aguinaldo, err1 := calcularAguinaldo(150000.00)

	if err != nil {
		fmt.Println(err)
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(1)
		} else {
			fmt.Printf("El aguinado es de; %.2f", aguinaldo)
			os.Exit(1)
		}
	} else {
		fmt.Printf("El salario del mes es: %.2f", salarioMensual)
		fmt.Println()
		if err1 != nil {
			fmt.Println(err1)
			os.Exit(1)
		} else {
			fmt.Printf("El aguinado es de: %.2f", aguinaldo)
			os.Exit(1)
		}
	}

}

func calcularSalarioMensual(minutos int) (float64, error) {
	horasTrabajadas := float64(minutos / 60)
	if horasTrabajadas < 80.00 {
		//return 0.00, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales y la cantidad de horas trabajadas fueron de: %.2f", horasTrabajadas)
		return 0.00, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		salarioBasePorHora := 1200.00
		salarioMensual := horasTrabajadas * salarioBasePorHora
		if salarioMensual >= 150000 {
			salarioMensualConImpuestoAplicado := salarioMensual * 0.9 //aplico el impuesto del 10%
			return salarioMensualConImpuestoAplicado, nil
		}
		return salarioMensual, nil
	}
}

func calcularAguinaldo(mejorSalario float64) (float64, error) {
	if mejorSalario < 0.00 {
		//return 0.00, fmt.Errorf("error: el mejor salario no puede ser 0 ni negativo y el valor que se esta pasando es: %.2f", mejorSalario)
		return 0.00, errors.New("error: el mejor salario no puede ser 0 ni negativo")
	} else {
		aguinaldo := mejorSalario / 2
		return aguinaldo, nil
	}
}
