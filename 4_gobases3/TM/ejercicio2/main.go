package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Product struct {
	id     string
	price  float64
	amount int
}

func main() {
	var option int
	var aux int = 1
	var products []Product
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Agregar producto al archivo\n2. Ver archivo \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			product := createProduct()
			products = append(products, product)

			var chain string = ""
			for i, p := range products {
				if i == 0 {
					chain += "ID;PRECIO;CANTIDAD"
				}
				chain += "\n" + p.id + ";" + fmt.Sprint(p.price) + ";" + fmt.Sprint(p.amount)
			}
			data := []byte(chain)

			ioutil.WriteFile("./archivo.txt", data, 0644)
		case 2:
			watchFile()
		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func createProduct() Product {
	var id string
	var price float64
	var amount int

	fmt.Println("Ingrese el id del producto")
	fmt.Scan(&id)

	fmt.Println("Ingrese el precio del producto")
	fmt.Scan(&price)

	fmt.Println("Ingrese la cantidad de producto")
	fmt.Scan(&amount)

	product := Product{id, price, amount}
	return product
}

func watchFile() {
	var total float64

	dat, err := ioutil.ReadFile("./archivo.txt")
	if err == nil {
		lines := strings.Split(string(dat), "\n")
		for _, line := range lines {
			values := strings.Split(line, ";")
			price, _ := strconv.ParseFloat(values[1], 64)
			amount, _ := strconv.ParseFloat(values[2], 64)
			total += price * amount
			fmt.Printf("%s\t\t\t%s\t%s\n", values[0], values[1], values[2])
		}
		fmt.Printf("\t\t\t%.2f\t\n", total)

	}

}
