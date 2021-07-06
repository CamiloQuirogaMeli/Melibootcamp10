package main

import "fmt"

func main() {
	var numero int = 10

	if numero > 0 && numero < 13 {
		meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
		fmt.Println(meses[numero-1])
	} else {
		fmt.Println("Mes incorrecto")
	}
}

//Se podria realizar con un Switch, Slice o con Maps
//Elijo hacer un Array.
