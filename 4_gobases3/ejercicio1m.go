package main

import (
	"fmt"
	"io/ioutil"
)

// Una empresa que se encarga de vender productos de limpieza necesita:
// 1. Implementar una funcionalidad para guardar un archivo de texto, con la información
// de productos comprados, separados por punto y coma.
// 2. Debe tener el id del producto, precio y la cantidad.
// 3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

type Product struct {
	id       string
	price    float64
	quantity int
}

func saveFile() {
	// prod1 := Product{"1", 100.2, 1}
	// prod2 := Product{"2", 250, 10}
	// prod3 := Product{"3", 500, 5}

	// d1 := []byte(string(prod1.id + "\n"))
	d2 := []byte("1 1 100.2; 2 5 250; 3 2 500\n")

	ioutil.WriteFile("./archivo.txt", d2, 0644)

	fmt.Println("Proceso Finalizado...")
}
