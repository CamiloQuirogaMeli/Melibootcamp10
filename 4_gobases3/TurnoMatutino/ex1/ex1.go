package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	permissonMode = 0777
	path          = "./file.json"
)

type Product struct {
	Id       int     `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func main() {

	values := generateMock()

	if values != nil {

		result := saveFile(values)

		if result == nil {
			fmt.Println("file created successfully")
		} else {
			fmt.Println("something went wrong on save file")
		}
	} else {
		fmt.Println("something went wrong on marshal")
	}
}

func generateMock() []byte {

	p := []Product{
		{1, 300.0, 500},
		{2, 650.0, 200},
		{3, 1000.0, 3},
	}

	if blob, err := json.Marshal(p); err == nil {
		return blob
	} else {
		return nil
	}
}

func saveFile(data []byte) error {
	return ioutil.WriteFile(path, data, permissonMode)
}

/*
Ejercicio 1 - Guardar archivo

Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma.
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/
