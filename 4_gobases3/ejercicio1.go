package main

import (
	"io/ioutil"
)

var archivo string = "111223,30012,1,444321,10000,4,434321,4030062,1,156845,32154,9"

func ejercicio1() {

	ioutil.WriteFile("archivo", []byte(archivo), 0644)

}
