package main

import (
	"io/ioutil"
)

func main() {
	productos := `{"id": 10, "precio": 500, "cantidad": 3};{"id": 8, "precio": 70, "cantidad": 25};{"id":15, "precio":56, "cantidad":6}`
	byteProductos := []byte(productos)
	ioutil.WriteFile("./productos.txt", byteProductos, 0644)
}
