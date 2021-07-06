package main

import "fmt"

func main() {
	// Ejercicio 1
	// var salary int = 140000
	// if salary < 150000 {
	// 	fmt.Println(MyError{"error: el salario ingresado no alcanza el mínimo imponible"})
	// } else {
	// 	fmt.Println("Debe pagar impuesto")
	// }

	// Ejercicio 2
	// var salary2 int = 140000
	// if salary2 < 150000 {
	// 	fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	// } else {
	// 	fmt.Println("Debe pagar impuesto")
	// }

	// Ejercicio 3
	var salary2 int = 140000
	err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary2)
	if salary2 < 150000 {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
