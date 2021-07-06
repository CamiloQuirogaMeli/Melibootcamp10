package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
Ejercicio 1 - Guardar archivo

Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma.
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/

/*Defino estructura producto*/
type producto struct {
	idProducto string
	precio     string
	cant       string
}

/*guarda un archivo con información (id precio y cantidad)sobre productos comprados,
separados por punto y coma*/
func guardaArchivo(productos []producto, filename string) {

	var text string

	/*para no perder la información del archivo*/
	contenido, errLectura := ioutil.ReadFile(filename)
	text += string(contenido)
	fmt.Println("Error de lectura: ", errLectura)

	/*Genero los productos a agregar con el formato solicitado*/

	for i := 0; i < len(productos); i++ {
		if i == len(productos)-1 {
			text += fmt.Sprintln(productos[i].idProducto, "-", productos[i].precio, "-",
				productos[i].cant)
		} else {
			text += fmt.Sprintln(productos[i].idProducto, "-", productos[i].precio, "-",
				productos[i].cant+";")
		}
	}

	/*Escribo en el archivo*/
	textGuardar := []byte(text)
	errEscritura := ioutil.WriteFile(filename, textGuardar, 0644)
	fmt.Println("Contenido archivo\n", string(textGuardar))
	fmt.Println("Error de lectura: ", errEscritura)

}

func main() {

	/*archivo donde se guardaran los productos*/
	filename := "/Users/areinhardt/Documents/4_basesGO-TM/TM-ejercicio1/prodLimpieza.txt"

	/*Busco el archivo*/
	f, err := os.Stat(filename)

	/*Detalle del archivo*/
	fmt.Println("Es un directorio: ", f.IsDir())
	fmt.Println("Nombre: ", f.Name())
	fmt.Println("Tamaño en Bytes: ", f.Size())
	fmt.Println("Fecha y hora de modificación: ", f.ModTime())
	fmt.Println("Permisos de archivos: ", f.Mode())

	fmt.Println("Error: ", err)

	p1 := producto{idProducto: "1", precio: "45.30", cant: "100"}
	p2 := producto{idProducto: "2", precio: "298.20", cant: "10"}
	p3 := producto{idProducto: "3", precio: "3789.00", cant: "200"}
	p4 := producto{idProducto: "4", precio: "1028.0", cant: "5"}

	var compra []producto
	compra = append(compra, p1, p2, p3, p4)

	guardaArchivo(compra, filename)

}
