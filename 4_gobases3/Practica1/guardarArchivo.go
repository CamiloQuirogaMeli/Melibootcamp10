package main

import "io/ioutil"
import "fmt"

//Una empresa que se encarga de vender productos de limpieza necesita:
//1. Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma.
//2. Debe tener el id del producto, precio y la cantidad.
//3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

func main() {
	text := []byte("idProducto;precio;cantidad\n1;1000;3\n2;650;2\n3;5000;1")
	err := ioutil.WriteFile("productos.csv", text, 0644)
    if err != nil {
        fmt.Println("ERROR:", err)
    }
}