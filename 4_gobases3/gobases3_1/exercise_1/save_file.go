package main

import (
	"fmt"
	f "fmt"
	"io/ioutil"
)

type Product struct {
	id     string
	price  string
	amount string
}

func (p *Product) newProduct() {

	f.Print("Enter the product id: ")
	f.Scanln(&p.id)

	f.Print("Enter the product price: ")
	f.Scanln(&p.price)

	f.Print("Enter the product amount: ")
	f.Scanln(&p.amount)
}

func addProduct() []Product {
	var products []Product
	i := 1
	for i == 1 {
		product := Product{}
		product.newProduct()
		products = append(products, product)
		fmt.Print("You want to add another product? (yes 1 : no 2): ")
		fmt.Scanln(&i)
	}
	return products
}

func saveFile(products []Product) {
	var msg string
	for i, val := range products {
		msg += val.id + "," + val.price + "," + val.amount
		if len(products) > (i + 1) {
			msg += ";"
		}
	}
	data := []byte(msg)
	err := ioutil.WriteFile("gobases3_1/save_file.txt", data, 0644)
	if err != nil {
		f.Println("Error saving file: ", err)
	} else {
		f.Println("File saved successfully")
	}
}

func main() {
	products := addProduct()
	saveFile(products)
}
