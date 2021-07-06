package main

func main() {
	imprimirResultados(calcularPromedio(2, 2, 3, 4, 5, 6, 7, 8, 10, 9, 8, 9, 8, 9, 7))
	imprimirResultados(calcularPromedio(2, 3, 4))
	imprimirResultados(calcularPromedio(-2, 3, 4, 5))
}

// func calcularPromedio(notas ...int) (promedio float64, err error) {
// 	var sumaTotal int
// 	for _, valor := range notas {
// 		if valor < 0 {
// 			err = errors.New("ningún valor puede ser menor a 0")
// 		} else {
// 			sumaTotal += valor
// 		}
// 	}

// 	promedio = float64(sumaTotal) / float64(len(notas))

// 	return
// }

// func imprimirResultados(promedio float64, err error) {
// 	if err == nil {
// 		fmt.Println("El promedio es: ", promedio)
// 	} else {
// 		fmt.Println(err)
// 	}
// }
