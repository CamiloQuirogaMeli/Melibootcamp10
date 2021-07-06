package main
import (
	"fmt"
	"io/ioutil"
	"strings"
)

/* Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)

Ejemplo:
ID Precio Cantidad
111223 30012.00 1
444321 1000000.00 4
434321 50.50 1 */

func main() {
	text, err := ioutil.ReadFile("../1/productosDeLimpieza.txt")

	if err != nil {
		fmt.Println("Error")
	} else {
		textParaseado := strings.ReplaceAll(string(text), ",", "\t")
		fmt.Printf("ID\tPrecio\tCantidad\n%s", textParaseado)
	}
}
