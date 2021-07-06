package main

import "fmt"

func main() {
	var mes int
	var calendario = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	fmt.Println("Ingrese el numero de mes (del 1 al 12)")
	fmt.Scanln(&mes)
	if !(0 < mes && mes < 13) {
		fmt.Println("Error, el nuemero ingresado no corresponde a nigun mes")
	} else {
		fmt.Println("El mes correspondiente es: ", calendario[mes])
	}
}
