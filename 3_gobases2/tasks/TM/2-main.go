package main

import "fmt"

func main() {
	qList := []float64{3, 4, 3.4, 4.3, 2.7}

	media, err := calcularPromedio(qList)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("El promedio de las notas es:", media)
}
