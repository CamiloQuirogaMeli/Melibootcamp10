package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity int
}

type Service struct {
	name    string
	price   float64
	minutes int
}

type Maintenance struct {
	name  string
	price float64
}

func sumProducts(products []Product, c chan float64) {
	var total float64
	for _, product := range products {
		total += product.price * float64(product.quantity)
	}
	c <- total
}

func sumServices(services []Service, c chan float64) {
	var total float64

	for _, service := range services {
		if service.minutes < 30 {
			total += service.price
		} else {
			total += service.price * (float64(service.minutes) / 30)
		}
	}
	c <- total
}

func sumMaintenance(maintenances []Maintenance, c chan float64) {
	var total float64
	for _, man := range maintenances {
		total += man.price
	}
	c <- total
}

func main() {
	var products = []Product{
		{name: "telefono", price: 1000, quantity: 2},
		{name: "radio", price: 100, quantity: 6},
	}
	var services = []Service{
		{name: "atencion de llamadas", price: 100, minutes: 20},
		{name: "soporte tecnico", price: 200, minutes: 60},
	}

	var maintenances = []Maintenance{
		{name: "mantenimiento a telefonos", price: 300},
		{name: "mantenimiento a radios", price: 600},
	}
	var chanProd = make(chan float64)
	var chanService = make(chan float64)
	var chanMant = make(chan float64)

	go sumProducts(products, chanProd)
	go sumServices(services, chanService)
	go sumMaintenance(maintenances, chanMant)

	var totalProducts float64 = <-chanProd
	fmt.Println("Total en productos:", totalProducts)

	var totalServicios float64 = <-chanService
	fmt.Println("Total en servicios:", totalServicios)

	var totalMant float64 = <-chanMant
	fmt.Println("Total en mantenimiento:", totalMant)

	var total = totalProducts + totalServicios + totalMant
	fmt.Println("Total: ", total)

}
