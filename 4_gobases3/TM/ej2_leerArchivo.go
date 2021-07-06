package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)


//.ReadFile() --> para leer archivo
dat, err := ioutil.ReadFile("./miArchivo.txt")
fmt.Println(dat, err)
*/

func formatearData(data string) []Producto {
	var formatData []Producto
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		token := strings.Split(line, " ; ")
		if len(token) != 0 {
			producto, err := leerProducto(token)
			if err == nil {
				formatData = append(formatData, producto)
			}
		}
	}
	return formatData
}

func leerProducto(args []string) (Producto, error) {
	product := Producto{}

	theId, err := strconv.Atoi(args[0])
	if err != nil {
		return product, err
	} else {
		product.id = theId
	}
	/*
		theName, err := args[1]
		if err != nil {
			return product, err
		} else {
			product.nombre = theName
		}

	*/
	theCant, err := strconv.Atoi(args[3])
	if err != nil {
		return product, err
	} else {
		product.cantidad = theCant
	}

	thePrice, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return product, err
	} else {
		product.precio = thePrice
	}

	return product, nil
}

func printFormatData(products []Producto) {
	fmt.Println("ID\t\tPrecio\tCantidad")
	acumPrecio := 0
	for _, product := range products {
		fmt.Print(product.id, "\t\t", product.precio, "\t", product.cantidad, "\nAcumulado")
		acumPrecio += int(product.precio) * product.cantidad
		fmt.Print("\t", acumPrecio, "\t\n")
	}

}
