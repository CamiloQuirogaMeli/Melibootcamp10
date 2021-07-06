package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Product struct {
	id       string
	price    float64
	quantity int
}

func saveProducts(products []Product) {
	var data string
	for _, product := range products {
		productLine := fmt.Sprintf("%s;%.2f;%d\n", product.id, product.price, product.quantity)
		data += productLine
	}
	d1 := []byte(data)
	err := ioutil.WriteFile("./products.txt", d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done. File saved correctly")
}

func main() {
	var products []Product
	product1 := Product{id: "df2sase122", price: 2900.99, quantity: 5}
	product2 := Product{id: "judsa76732", price: 3000, quantity: 2}
	product3 := Product{id: "x87a6x4z31", price: 5000, quantity: 4}
	products = append(products, product1)
	products = append(products, product2)
	products = append(products, product3)
	saveProducts(products)

}
