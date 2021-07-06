package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

func ejercicio9() {

	dataFromFile, errorFromReadingFile := ioutil.ReadFile("./dataJul-5-2021.json")

	var productos []Producto
	flete := Flete{
		maxCantidad: 10000000,
	}

	if errorFromReadingFile == nil {
		if errorFromUnmarshal := json.Unmarshal(dataFromFile, &productos); errorFromUnmarshal == nil {

			for _, producto := range productos {
				flete.agregarProducto(producto)
			}

		}
	}

	fmt.Printf("Envios totales: %v\n", flete.calcularCantidadEnvios())

}

const (
	CHICO    = "chico"
	FRAGIL   = "fragil"
	ESPECIAL = "especial"
)

var adicional = map[string]float64{CHICO: 1, MEDIANO: 1.05, GRANDE: 1.2, FRAGIL: 1.75, ESPECIAL: 1}

type Producto struct {
	Tipo   string
	Tamano float64
}

func (producto *Producto) tamanoTotal() float64 {
	return producto.Tamano * adicional[producto.Tipo]
}

type Flete struct {
	productos   []Producto
	maxCantidad float64
}

func (flete *Flete) agregarProducto(producto Producto) {
	flete.productos = append(flete.productos, producto)
}

func (flete *Flete) calcularCantidadEnvios() float64 {

	var cantidadActualNormal, cantidadActualEspecial, countNormal, countEspecial float64

	for _, producto := range flete.productos {

		if producto.Tipo != ESPECIAL {
			cantidadActualNormal += producto.tamanoTotal()
			countNormal++
		} else {
			cantidadActualEspecial += producto.tamanoTotal()
			countEspecial++
		}

	}

	var enviosNormales float64 = math.Ceil(cantidadActualNormal / flete.maxCantidad)

	fmt.Printf("Envios normales: %v\nCantidad productos: %v\nTamano total: %.2f cm3\n", enviosNormales, countNormal, cantidadActualNormal)

	var enviosEspeciales float64 = math.Ceil(cantidadActualEspecial / flete.maxCantidad)

	fmt.Printf("Envios especiales: %v\nCantidad productos: %v\ntamano total: %.2f cm3\n", enviosEspeciales, countEspecial, cantidadActualEspecial)

	return enviosNormales + enviosEspeciales
}
