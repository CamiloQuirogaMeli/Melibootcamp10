package main

import (
	"errors"
	"fmt"
)

const (
	MINIMO float64 = 50000
	ALTO   float64 = 150000
)

func main() {
	var seguir bool = true

	for seguir {

		fmt.Println("Ingresá el sueldo: ")
		var sueldo float64
		fmt.Scanln(&sueldo)

		p, i, sn, err := calcularImpuesto(sueldo)

		if err == nil {
			fmt.Println("Porcentaje de impuesto: ", p, "%")
			fmt.Println("Monto impuesto: ", i)
			fmt.Println("Sueldo neto: ", sn)
		} else {
			fmt.Println(err)
		}

		fmt.Println("¿Querés calcular otro impuesto? Ingresa true o false")
		fmt.Scanln(&seguir)

	}

}

func calcularImpuesto(sueldo float64) (porcentaje float64, impuesto float64,
	sueldoNeto float64, err error) {

	switch {
	case sueldo < MINIMO && sueldo > 0:
		porcentaje = 0
	case sueldo > MINIMO && sueldo < ALTO:
		porcentaje = 17
	case sueldo > ALTO:
		porcentaje = 27
	case sueldo <= 0:
		err = errors.New("el valor ingresado es no válido")
	}

	impuesto = sueldo * porcentaje / 100
	sueldoNeto = sueldo - impuesto

	return
}
