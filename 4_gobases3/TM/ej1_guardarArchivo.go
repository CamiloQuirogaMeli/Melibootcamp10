package main

import (
	"fmt"
	"io/ioutil"
)

/*
Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma.
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

/*
//.WriteFile() --> para escribir archivo
d1 := []byte("hello\ngo\n")
//err := ioutil.WriteFile("./dat1", d1, 0644)
fmt.Println(d1)
*/
type Producto struct {
	nombre   string
	id       int
	precio   float64
	cantidad int
}

var nextId int = 1

func nuevoProducto(nombre string, id int, precio float64, cant int) Producto {
	prod := Producto{nombre: nombre, id: id, precio: precio, cantidad: cant}
	prod.id = nextId
	nextId++
	return prod
}

func (p Producto) toString() string {
	return fmt.Sprint(p.id, " ; ", p.nombre, " ; ", p.precio, " ; ", p.cantidad)
}

func guardarArchivo(p []Producto) {
	toWrite := ""
	for i, value := range p {
		toWrite += fmt.Sprint(value.toString())
		if i != len(p)-1 {
			toWrite += "\n"
		}
	}
	d := []byte(toWrite)
	err := ioutil.WriteFile("./lista_productos.txt", d, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Exito")
	}

}
