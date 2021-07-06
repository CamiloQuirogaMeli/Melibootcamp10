package main

import "fmt"
import "errors"

//Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
//y nos devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

func main()  {
	fmt.Println("Vamos a calcular el promedio para el estudiante 1, 2 y 3...")

	fmt.Print("El estudiante 1 tiene como promedio:")
	val, err := promedio(3.4, 4.2, 5, 3.7, 1.5, 1.9, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	fmt.Println("El estudiante 2 tiene como promedio:")
	val1, err1 := promedio(5, 5, 4.5, 4.3, 5, 4.8, 3)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(val1)
	}

	fmt.Println("El estudiante 3 tiene como promedio:")
	val2, err2 := promedio(5, 5, -4.5, 4.3, -3)

	if err2 != nil {
		fmt.Println("ERROR:", err2)
	} else {
		fmt.Println(val2)
	}
}

func promedio(notas ...float64) (float64, error){
	var prom float64
	fmt.Println("Las notas son:", notas)

	for _, value := range notas{
		if value < 0 {
			return 0, errors.New("Uno de los números ingresados es negativo.")
		} 

		prom += value
	}

	return prom / float64(len(notas)), nil
}
