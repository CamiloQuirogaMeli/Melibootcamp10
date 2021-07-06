package main

import "fmt"
import "encoding/json"

func main() {

	fmt.Println("* Caso crear un nuevo product")
	product := NewProduct("Cereal", 16.5)
	fmt.Println("\tDetalles del producto creado:")
	productJson, _ := json.MarshalIndent(product, "\t", " ")
	fmt.Println("\t",string(productJson))

	fmt.Println("* Caso agregar producto a usuario")
	user := User{}
	AddProductsToUser(&user, &product, 5)
	fmt.Println("\tDetalles del carrito del usuario luego de agregar:")
	userCartJson, _ := json.MarshalIndent(user, "\t", " ")
	fmt.Println("\t",string(userCartJson))

	fmt.Println("* Caso remover productos del usuario")
	RemoveProductsFromUser(&user)
	fmt.Println("\tDetalles del carrito del usuario luego de eliminar todos los productos:")
	userCartJson, _ = json.MarshalIndent(user, "\t", " ")
	fmt.Println("\t",string(userCartJson))
}