package main

import (
	"errors"
	"fmt"
)

type errorSalario struct {
}

func (e errorSalario) Error() string{
	return "error: el trabajador no puede haber trabajado menos de 80 hs mensuales"
}

func calcSalarioMensual(horasT float64, valorHs float64) (float64, error){

	if horasT < 80{
		e2 := errorSalario{}
		e1 := fmt.Errorf("error al calcular salario: %w", e2)
		return 0, e1
	}

	salario := horasT * valorHs

	if salario >= 150000 {
		salario = impuestoSalario(salario)
	}
	return salario, nil
}

func impuestoSalario (sal float64) float64 {
	return sal * 0.90
}

func calcMedioAguinaldo (salario[6] float64) (float64){
	var maxSalario float64 = 0
	for i:=0; i<len(salario); i++{

		if salario[i]<0{
			fmt.Println(errors.New("no se puede ingresar un salario negativo para calcular aguinaldo"))
			return 0
		}
		if salario[i]>maxSalario{
			maxSalario = salario[i]
		}
	}
	aguinaldo := (maxSalario / 12)*float64(len(salario))

	return aguinaldo
}

func main (){
	var valorHora, horasTrabajadas float64 = 2000, 9

	var salarioSemestre[6] float64
	salarioSemestre[0] = 131230
	salarioSemestre[1] = 181230
	salarioSemestre[2] = 141430
	salarioSemestre[3] = 150000.8
	salarioSemestre[4] = 167230.45
	salarioSemestre[5] = 190230

	salario, err := calcSalarioMensual(horasTrabajadas, valorHora)

	if err != nil {
		fmt.Println(errors.Unwrap(err))
	}else{
		fmt.Println("El salario del trabajador es de: $", salario)
	}

	aguinaldo := calcMedioAguinaldo(salarioSemestre)

	if aguinaldo != 0{
		fmt.Println("El aguinaldo que le corresponde al trabajador es: $", aguinaldo)
	}
}

/*Vamos a hacer que nuestro programa sea un poco más complejo y útil.
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
retornos de error en tu función “main()”.*/