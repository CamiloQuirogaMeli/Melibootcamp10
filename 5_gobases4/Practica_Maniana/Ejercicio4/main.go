package main

import (
	"errors"
	"fmt"
)

func salarioMensual(valorHora, cantHoras, impuesto, minhoras float64) (float64, error) {
	var salario float64
	if cantHoras < minhoras {
		err := fmt.Errorf("error: el trabajador no puede haber trabajado menos de %.2f hs mensuales", minhoras)
		return 0, errors.New(fmt.Sprint(err))

	}
	if (valorHora * cantHoras) >= impuesto {
		salario = (valorHora * cantHoras) - ((valorHora * cantHoras) * 0.10)
	} else {
		salario = (valorHora * cantHoras)
	}
	return salario, nil
}

func medioAguinaldo(mejorSalario float64, cantMeses int) (float64, error) {
	var monto float64
	if mejorSalario < 0 || cantMeses < 0 {
		return 0, errors.New("error: no debe ingresar numeros negativos")
	}
	monto = (mejorSalario / 12) * float64(cantMeses)
	return monto, nil
}

func main() {
	salario, err1 := salarioMensual(500, 80, 150000, 80)
	aguinaldo, err2 := medioAguinaldo(100000, 6)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Print("Salario: ", salario, "\n")
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Print("Medio Aguinaldo: ", aguinaldo)
	}

}
