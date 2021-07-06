package main

import (
	"io/ioutil"
)

func main() {
	productos := []byte("Id:1 Precio:5 Cantidad: 2; Id:2 Precio:3 Cantidad: 10; Id:13 Precio:16 Cantidad: 20;")
	ioutil.WriteFile("productos.txt", productos, 0644)
}
