package main

import (
	"fmt"
)

func main() {

	mes := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	var numero int = 11 //elija y cambie este numero

	fmt.Println("El mes que eligio fue el: ", numero, mes[numero-1])

}
