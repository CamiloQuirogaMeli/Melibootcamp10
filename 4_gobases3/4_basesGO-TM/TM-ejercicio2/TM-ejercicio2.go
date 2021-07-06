package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*

La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el
ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se
debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

*/

func leeArchivo(filename string) string {

	text, errLectura := ioutil.ReadFile(filename)
	fmt.Println("Error de lectura: ", errLectura)
	return string(text)

}

func main() {

	filename := "/Users/areinhardt/Documents/4_basesGO-TM/TM-ejercicio1/prodLimpieza.txt"
	contenido := leeArchivo(filename)

	productos := strings.Split(contenido, ";")

	fmt.Println("ID \t \t \t \t \t PRECIO \t \t CANTIDAD")

	total := 0.0

	for _, producto := range productos {
		datos := strings.Split(producto, " - ")
		precio, _ := strconv.ParseFloat(datos[1], 64)
		cant, _ := strconv.ParseFloat(datos[2], 64)
		total += precio * cant
		fmt.Println(datos[0], "\t\t\t\t\t", datos[1], "\t\t", datos[2])

	}

	fmt.Println("TOTAL\t\t\t\t\t", total)

}
