package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		Ejercicio 1 - Estructura un JSON
		Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
		Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
		Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
		Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
		1. Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
		2. Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.
	*/

	//CREACION DE SLICE DE USUARIOS EN JSON
	u1 := User{001, "Martina", "Olivera", "martina@correo.com", 23, 163, true, "02-07-2021"}
	u2 := User{002, "Pepito", "Perez", "pepito@correo.com", 33, 170, true, "02-07-2021"}
	u3 := User{003, "Pepito", "Perez", "pepito@correo.com", 33, 170, true, "02-07-2021"}
	u4 := User{004, "Pepito", "Perez", "pepito@correo.com", 33, 170, true, "02-07-2021"}
	u5 := User{005, "Pepito", "Perez", "pepito@correo.com", 33, 170, true, "02-07-2021"}

	us := []User{}
	us = append(us, u1, u2, u3, u4, u5)
	jsonDate, err := json.Marshal(us)

	if err != nil {
		fmt.Println(err)
	}

	err1 := ioutil.WriteFile("users.json", jsonDate, 0644)

	if err1 != nil {
		fmt.Println(err1)
	}

}

type User struct {
	Id           int
	Name         string
	Surname      string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate string
}
