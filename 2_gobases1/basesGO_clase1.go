package main

import (
	"fmt"
	"time"
)

func main() {

	//arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0])
	fmt.Println(a[0], a[1])
	//fmt.Printf(a[0], a[1]) //debo aclarar el tipo de dato para usar el printf
	fmt.Println(a)

	//slices
	var s = []bool{true, false}
	fmt.Println(s[0])

	sli := make([]int, 5)
	fmt.Println(sli)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])

	var sliceI []int
	sliceI = append(sliceI, 2, 3, 4)
	fmt.Println(sliceI)

	//MAPS
	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}

	fmt.Println(students)             //da el mapa completo
	fmt.Println(students["Benjamin"]) //da el valor del ìndice que correpsonde a Benjamin
	//fmt.Println(students[20]) DA ERROR

	//agregar valores
	students["Diego"] = 10
	students["Soledad"] = 5
	students["Ariel"] = 40
	fmt.Println(students)

	//modificar valor
	students["Soledad"] = 0
	fmt.Println(students)

	//borrar elemento
	delete(students, "Benjamin")
	fmt.Println(students)

	//recorrer elementos
	for key, element := range students {
		fmt.Println("Key: ", key, "=>", "Element: ", element)
	}

	//IF ... ELSE
	var sueldo float64 = 5000
	if sueldo <= 3000 {
		fmt.Println(" Esta persona debe pagar impuestos")
	} else if sueldo <= 4000 {
		fmt.Println(" Debe pagar $%4.2f de su sueldo\n", (sueldo/100)*10)
	} else {
		fmt.Println(" Debe pagar $%4.2f de su sueldo\n", (sueldo/100)*15)
	}

	//SWITCH
	day := 1
	switch day {
	case 0:
		fmt.Println("Lunes")
	case 1:
		fmt.Println("Martes")
	case 2:
		fmt.Println("Miercoles")
	default:
		fmt.Println("Desconocido")

	}

	today := time.Now()
	var t int = today.Day()
	switch t {
	case 5, 10, 15:
		fmt.Println("Comprar comida")
	case 1, 2, 3, 4:
		fmt.Println("Limpia tu casa")
		fallthrough
	case 25:
		fmt.Println("No hay fiesta")
	default:
		fmt.Println("No hay info disponible")
	}

	//FOR  para todos
	frutas := []string{"manzana", "banana", "pera"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}
}
