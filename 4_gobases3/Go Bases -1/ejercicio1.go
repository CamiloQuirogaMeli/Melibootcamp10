package main

import (
	"fmt"
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

var path = "./productos.txt"

func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado correctamente, en", path)
}

func escribirArchivo() {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	id := "1"
	precio := "500"
	cantidad := "2"
	_, err = file.WriteString("id: ")
	_, err = file.WriteString(id)
	_, err = file.WriteString(",")
	_, err = file.WriteString("precio: ")
	_, err = file.WriteString(precio)
	_, err = file.WriteString(",")
	_, err = file.WriteString("cantidad: ")
	_, err = file.WriteString(cantidad)
	_, err = file.WriteString(",")
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	crearArchivo()
	escribirArchivo()
}
