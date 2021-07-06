package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Ejercicio 2 - Leer archivo
	La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
	por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
	y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
	visualizar el total (Sumando precio por cantidad)
*/

var path = "./productos2.txt"

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

	_, err = file.WriteString("ID        PRECIO       CANTIDAD\n")
	_, err = file.WriteString("111221    30012.00     1\n")
	_, err = file.WriteString("444321    1000000.00   4\n")
	_, err = file.WriteString("434321    50.50        1\n")
	_, err = file.WriteString("          4030062.50     \n")

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

func leerArchivo() {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//fmt.Println(scanner.Bytes())
	}
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
	leerArchivo()
}
