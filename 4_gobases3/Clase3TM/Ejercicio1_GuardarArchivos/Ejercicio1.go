// Una empresa que se encarga de vender productos de limpieza necesita:
// 1. Implementar una funcionalidad para guardar un archivo de texto, con la información
// de productos comprados, separados por punto y coma.
// 2. Debe tener el id del producto, precio y la cantidad.
// 3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	"fmt"
	"io/ioutil"
)

func saveFile(text []byte) {
	err := ioutil.WriteFile("./pruebaGuardado.txt", text, 0644)
	if err != nil {
		fmt.Printf("En teoria quedo bien..")
	}
}

func main() {

	var productName string = "Detergente"
	//var id int = 4444
	var precio float64 = 45.66
	var cantidad float64 = 6
	str := fmt.Sprintln(productName, ";", precio, ";", cantidad, "\n")
	d1 := []byte(str)
	saveFile(d1)

}
