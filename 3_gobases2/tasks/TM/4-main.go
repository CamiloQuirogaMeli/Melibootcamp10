package main

import "fmt"

func main() {
	minimumFunc, errMin := calcularEstadisticas("minimo")
	maximumFunc, errMax := calcularEstadisticas("maximo")
	mediaFunc, errMedia := calcularEstadisticas("promedio")
	_, errError := calcularEstadisticas("error")

	if errMin != nil {
		fmt.Println(errMin)
	}

	if errMax != nil {
		fmt.Println(errMax)
	}

	if errMedia != nil {
		fmt.Println(errMedia)
	}

	if errError != nil {
		fmt.Println(errError)
	}

	minRes, errMinRes := minimumFunc(4.6, 2.9, 3.6, 4.9, 3.0, 4.2)
	maxRes, errMaxRes := maximumFunc(4.6, 2.9, 3.6, 4.9, 3.0, 4.2)
	mediaRes, errMediaRes := mediaFunc(4.6, 2.9, 3.6, 4.9, 3.0, 4.2)

	if errMaxRes != nil {
		fmt.Println(errMaxRes)
	}

	if errMinRes != nil {
		fmt.Println(errMinRes)
	}

	if errMediaRes != nil {
		fmt.Println(errMediaRes)
	}

	fmt.Println(minRes)
	fmt.Println(maxRes)
	fmt.Println(mediaRes)
}
