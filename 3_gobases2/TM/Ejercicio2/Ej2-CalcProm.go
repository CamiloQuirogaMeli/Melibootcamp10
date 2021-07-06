package main

import (
	"errors"
	"fmt"
)

func calcProm(notas ...float32) (float32, error){
	var suma float32 = 0

	for _, value := range notas{
		if value < 0{
			return 0, errors.New("Numero negativo ingresado")
		}else{
			suma += value
		}
	}

	return suma/float32(len(notas)), nil
}

func main(){
	var nota1, nota2, nota3, nota4 float32 = 5, 7.5, 9, 5.5

	promedio, err := calcProm(nota1, nota2, nota3, nota4)

	if err != nil{
		fmt.Println("Se ingresó un numero negativo, pruebe nuevamente.")
	}else{
		fmt.Println("El promedio del alumno es: ", promedio)
	}
}

/* Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo */