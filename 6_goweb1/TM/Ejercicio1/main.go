package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type product struct {
	Id        int
	Nombre    string
	Precio    int
	Publicado bool
	Color     string
	Fecha     string
	Stock     int
	Codigo    string
}

type usuario struct {
	Id       int
	Nombre   string
	Apellido string
	Email    string
	Edad     int
	Altura   int
	Activo   bool
	Fecha    string
}

type transaccion struct {
	Id       int
	Codigo   string
	Moneda   string
	Monto    float64
	Emisor   string
	Receptor string
	Fecha    string
}

func main() {

	p := product{
		Nombre:    "MacBook Pro",
		Precio:    1500,
		Publicado: true,
		Id:        150987,
		Color:     "gris",
		Codigo:    "123drftg",
		Stock:     750,
	}

	p1 := product{
		Nombre:    "Lenovo Pro",
		Precio:    1100,
		Publicado: true,
		Id:        150555,
		Color:     "negra",
		Codigo:    "654kijuhy",
		Stock:     701,
	}

	p2 := product{
		Nombre:    "Dell Pro",
		Precio:    1110,
		Publicado: true,
		Id:        852147,
		Color:     "gris",
		Codigo:    "987ploi",
		Stock:     520,
	}

	productos := []product{p, p1, p2}

	productos1 := []product{}

	jsonData, err := json.Marshal(productos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	err2 := ioutil.WriteFile("productos.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err2)
	}

	fmt.Print(productos1)

	err1 := json.Unmarshal(jsonData, &productos1)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Print(productos1)
}
