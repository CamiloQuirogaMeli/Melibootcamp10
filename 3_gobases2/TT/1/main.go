package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	alumno := Alumno{
		Nombre:   "Lucas",
		Apellido: "Ficarra",
		DNI:      "40109428",
		Fecha:    "22/06/2021",
	}

	detalle, err := alumno.detalle()

	if err == nil {
		fmt.Println(detalle)
	} else {
		fmt.Println(err)
	}
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() (string, error) {
	miJSON, err := json.Marshal(a)

	return string(miJSON), err
}
