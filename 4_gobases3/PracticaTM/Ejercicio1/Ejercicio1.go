// Ejercicio 1 - Guardar archivo
// Una empresa que se encarga de vender productos de limpieza necesita:
// 1. Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma.
// 2. Debe tener el id del producto, precio y la cantidad.
// 3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	f "fmt"
	"os"
	s "strings"
)

type Products struct {
	idProd   int
	price    float64
	cantProd int
}

func writeFileProducts(fileName string, prod Products) {
	var buffer string

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		f.Println(err)
	}

	buffer = f.Sprintf("%d,%.2f,%d\n", prod.idProd, prod.price, prod.cantProd)
	defer file.Close()
	if _, err := file.WriteString(buffer); err != nil {
		f.Println(err)
	}
}

func main() {

	var prod Products
	var cantProdsAdd int
	var fileName string
	var decisionOtehrFile string = "si"

	for s.Compare(decisionOtehrFile, "si") == 0 {
		f.Println("Ingrese el nombre del archivo. Ejemplo: Produstos1")
		f.Println("Los archivos se guardaran en la carpeta \"Archivos\" que se encuentra en un directorio anterior. Se guardaran como archivos txt")
		f.Printf("Nombre del archivo: ")
		f.Scan(&fileName)
		fileName = "../Archivos/" + fileName + ".txt"
		_, err := os.Stat(fileName)
		for err == nil {
			f.Printf("El archivo ya existe, ingrese otro nombre: ")
			f.Scan(&fileName)
			_, err = os.Stat(fileName)
		}
		os.Create(fileName)
		f.Printf("¿Cuantos productos desea agregar?: ")
		f.Scan(&cantProdsAdd)
		for i := 0; i < cantProdsAdd; i++ {
			f.Printf("Ingrese el ID del producto %d: ", i+1)
			f.Scan(&prod.idProd)
			f.Printf("Ingrese el PRECIO del producto: ")
			f.Scan(&prod.price)
			f.Printf("Ingrese la CANTIDAD de productos: ")
			f.Scan(&prod.cantProd)
			writeFileProducts(fileName, prod)
		}
		f.Printf("¿Desea crear otro archivo?(si/no): ")
		f.Scan(&decisionOtehrFile)
		decisionOtehrFile = s.ToLower(decisionOtehrFile)
	}

}
