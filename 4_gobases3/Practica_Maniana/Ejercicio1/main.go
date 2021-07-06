package main

import (
	"fmt"
	"io/ioutil"
)

type Productos struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	p := Productos{1, 500, 2}
	txt := []byte(fmt.Sprint(p))
	ioutil.WriteFile("./archivoEj1", txt, 0644)
}
