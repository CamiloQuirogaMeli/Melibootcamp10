package main

import "errors"

const (
	topeImpuesto = 150000
)

func calcularSalarioMensual(horasTrabajadas float64) (float64, error) {
	valorHora := 1000.0
	var sueldo float64
	if horasTrabajadas < 80 {
		return 0, errors.New("no puede trabajar menos de 80 horas semanales")
	}
	sueldo = valorHora * horasTrabajadas
	if sueldo > topeImpuesto {
		sueldo = sueldo - (sueldo * 0.1)
	}
	return sueldo, nil
}

func calcularAguinaldo(mejorSalario float64, cantidadMesesTrabajados int) (float64, error) {
	if mejorSalario < 0 && cantidadMesesTrabajados < 0 {
		err := errors.New("no se pueden ingresar valores negativos")
		return 0, err
	} else {
		aguinaldo := (mejorSalario / 12) * float64(cantidadMesesTrabajados)
		return aguinaldo, nil
	}
}

func main() {
	horasTrabajadas := 30
	var sueldosMensuales []float64

	for i := 0; i < 6; i++ {
		sueldo, err := calcularSalarioMensual(float64(horasTrabajadas))
		if err != nil {
			sueldosMensuales = append(sueldosMensuales, sueldo)
		}
	}

	aguinaldo, err := calcularAguinaldo(sueldosMensuales[3], 6)

	if err != nil {
		println("El aguinaldo es de %f", aguinaldo)
	}

}
