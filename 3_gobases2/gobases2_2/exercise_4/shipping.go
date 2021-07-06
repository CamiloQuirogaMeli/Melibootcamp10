package main

import (
	er "errors"
	"fmt"
)

type Product struct {
	product_name string
	product_type string
	product_size float64
}

func (p *Product) newProduct() {

	fmt.Print("Enter the product name: ")
	fmt.Scanln(&p.product_name)

	fmt.Print("Enter the product size: ")
	fmt.Scanln(&p.product_size)

	i := 1
	for i == 1 {
		fmt.Print("Enter the product type: ")
		fmt.Scanln(&p.product_type)
		_, err := p.calculateTotalSize()
		i = handleErrors(err)
	}
}

const (
	Small   = "small"
	Medium  = "medium"
	Big     = "big"
	Fragile = "fragile"
	Special = "special"
)

func (p Product) calculateTotalSize() (float64, error) {
	var total_size float64
	switch p.product_type {
	case Small:
		total_size = p.product_size
	case Medium:
		total_size = p.product_size + p.product_size*0.05
	case Big:
		total_size = p.product_size + p.product_size*0.2
	case Fragile:
		total_size = p.product_size + p.product_size*0.75
	case Special:
		total_size = p.product_size
	default:
		return 0, er.New("\nThis type of product does not exist")
	}
	return total_size, nil
}

type Freight struct {
	products         []Product
	number_shipments int
}

func (f *Freight) addProduct() {
	i := 1
	for i == 1 {
		product := Product{}
		product.newProduct()
		total_size, _ := product.calculateTotalSize()
		println("The package requires a total size of", total_size, "cm^3")

		f.products = append(f.products, product)
		fmt.Print("You want to add another product? (yes 1 : no 2): ")
		fmt.Scanln(&i)

	}
}

func (f *Freight) calculateShipments() {
	sum := 0.
	f.number_shipments = 1
	for _, value := range f.products {
		total_size, _ := value.calculateTotalSize()
		sum += total_size
		if sum >= 10000000 {
			f.number_shipments++
			sum = 0.
		}
	}
}

func main() {
	freight := Freight{}
	freight.addProduct()
	freight.calculateShipments()
	fmt.Println("List of products to send:", freight.products)
	fmt.Println("They must be made", freight.number_shipments, "shipments")
}

func handleErrors(err error) int {
	if err == nil {
		return 2
	} else {
		fmt.Println(err)
		fmt.Println("Please try again adding another data ...")
		return 1
	}
}
