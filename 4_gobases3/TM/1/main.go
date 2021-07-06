package main
import (
	"fmt"
	"io/ioutil"
)

/* Ejercicio 1 - Guardar archivo

Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma.
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable. */

var products = [3]string{"1, 200, 2", "2, 1500, 1", "3, 2000, 5"}

func main() {
	var text []byte
	for _, val := range products {
		text = append(text, []byte(val+"\n")...)
	}
	err := ioutil.WriteFile("./productosDeLimpieza.txt", text, 0644)
	if err != nil {
		fmt.Println("Error")
	}
}