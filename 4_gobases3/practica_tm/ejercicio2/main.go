package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//
// Ejercicio 2 - Leer archivo
// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
// por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
// y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
// visualizar el total (Sumando precio por cantidad)
//

const (
	PATH = "../"
	FILE = "compras.txt"
)

func main() {

	fmt.Println("======= Leer archivo =======")
	// Mostramos el contenido del archivo
	content, err := ioutil.ReadFile(PATH + FILE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ID\tStock\tPrecio(Unidad)\tCantidad\tValor total\n")
		result := strings.ReplaceAll(string(content), ";", "\t")
		fmt.Printf("%s", result)
	}

}
