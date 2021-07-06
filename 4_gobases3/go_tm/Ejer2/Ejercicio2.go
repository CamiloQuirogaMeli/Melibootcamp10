/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima
por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID
y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	products, err := ioutil.ReadFile("../products.txt")
	if err != nil {
		fmt.Println("Algo fallo")
		os.Exit(0)
	}
	productsSeparate := strings.Split(string(products), ";")
	if len(productsSeparate) > 0 {
		productsSeparate = productsSeparate[:len(productsSeparate)-1]
	}
	//Regular expression to extract number from string
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	var result float64
	fmt.Println("\t ID \t \t Precio Cantidad")
	for _, product := range productsSeparate {
		//use regular expression
		submatchall := re.FindAllString(product, -1)
		fmt.Println("\t", submatchall[0], "\t \t", submatchall[1], "\t", submatchall[2])
		price, err := strconv.ParseFloat(submatchall[1], 64)
		quantity, err := strconv.ParseFloat(submatchall[2], 64)
		price = price * quantity
		if err != nil {
			fmt.Println("Error")
		}
		result += price
	}
	fmt.Println("\t \t \t", result)
}
