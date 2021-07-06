// Ejercicio 2 - Leer archivo
// La misma empresa necesita leer el archivo almacenado, para ello requiere que:
// se imprima por pantalla mostrando los valores tabulados,
// con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
// Ejemplo:
// ID Precio 	Cantidad
// 111223 		30012.00 1
// 444321 		1000000.00 4
// 434321 		50.50 1
// 4030062.50

package main

import (
	"bufio"
	f "fmt"
	"os"
	pf "path/filepath"
	"strconv"
	s "strings"
)

type Products struct {
	idProd   int
	price    float64
	cantProd int
}

func readFileProducts(optionFile int, files []string) {

	var line []string
	var prod Products

	file, err := os.Open(files[optionFile-1])
	if err != nil {
		f.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	f.Printf("%-v  %10v  %10v \n", "ID", "Precio", "Cantidad")
	for scanner.Scan() {
		line = s.Split(scanner.Text(), ",")
		prod.idProd, _ = strconv.Atoi(line[0])
		prod.price, _ = strconv.ParseFloat(line[1], 64)
		prod.cantProd, _ = strconv.Atoi(line[2])
		f.Printf("%-d %10v  %10v\n", prod.idProd, prod.price, prod.cantProd)
	}

}

func main() {

	var optionFile int

	files, err := pf.Glob("../Archivos/*.txt")
	if err != nil {
		f.Println(err)
		return
	}

	if len(files) == 0 {
		f.Println("No existen archivos para leer")
		return
	}

	f.Printf("¿Que archivo desea leer?\n")

	for i, file := range files {
		f.Printf("%d - %s\n", i+1, file)
	}

	f.Printf("Seleccione un número: ")
	f.Scan(&optionFile)

	for optionFile < 1 && optionFile > len(files) {
		f.Println("No elgió una opcion correcta")
		f.Printf("Ingrese nuevamente: ")
		f.Scan(&optionFile)
	}

	readFileProducts(optionFile, files)

}
