package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Product struct {
	Id        int
	Price     float64
	Inventory int
}

func main() {
	var products []Product
	write := false
	for !write {
		var menu string
		fmt.Print("¿Desea añadir un producto (1) o terminar y escribir el archivo (2)? (1/2): ")
		fmt.Scanln(&menu)
		menu = strings.ToLower(menu)
		if menu == "1" {
			var id int
			var price float64
			var inv int
			fmt.Print("Escriba el id del producto: ")
			fmt.Scanln(&id)
			fmt.Print("Escriba el precio del producto: ")
			fmt.Scanln(&price)
			fmt.Print("Escriba la cantidad de productos: ")
			fmt.Scanln(&inv)
			products = append(products, Product{id, price, inv})
		} else if menu == "2" {
			write = true
		}
	}
	data := ""
	for _, p := range products {
		data += fmt.Sprintf("%d;%.2f;%d\n", p.Id, p.Price, p.Inventory)
	}
	ioutil.WriteFile("../data.csv", []byte(data), 0644)
}
