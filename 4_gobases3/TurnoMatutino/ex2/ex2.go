package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	filePath = "./file2.json"
)

type Product struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {

	if jsonBlob, err := ioutil.ReadFile(filePath); err == nil {

		var products []Product

		if err := json.Unmarshal(jsonBlob, &products); err == nil {
			printValues(products)
		}
	}
}

func printValues(values []Product) {

	fmt.Printf("ID\tPRICE\tQUANTITY\n")
	var total float64

	for _, value := range values {
		fmt.Printf("%d\t%.2f\t%d\n", value.Id, value.Price, value.Quantity)
		total += value.Price
	}
	fmt.Printf("total: %.2f", total)
}

/*
Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)

Ejemplo:
ID Precio Cantidad
111223 30012.00 1
444321 1000000.00 4
434321 50.50 1

4030062.50
*/
