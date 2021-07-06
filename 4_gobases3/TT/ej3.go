package main

import (
	"fmt"
	"math/rand"
)

type Products struct {
	Name  string
	Price float64
	Cant  int
}
type Services struct {
	Name    string
	Price   float64
	Minutes int
}
type Maintenance struct {
	Name  string
	Price float64
}

func main() {
	cp := make(chan float64)
	cs := make(chan float64)
	cm := make(chan float64)
	products := createProducts(2)
	services := createServices(2)
	maintenance := createMaintenances(2)
	fmt.Println(products)
	fmt.Println(services)
	fmt.Println(maintenance)
	go totalProducts(products, cp)
	go totalServices(services, cs)
	go totalMaintenance(maintenance, cm)
	sum := <-cp + <-cs + <-cm
	fmt.Println("Total: ", sum)
}

func totalProducts(allP []Products, cp chan float64) {
	fmt.Println("products")
	var subTotal float64
	for _, value := range allP {
		subTotal += value.Price * float64(value.Cant)
	}
	cp <- subTotal
}

func totalServices(allS []Services, cs chan float64) {
	fmt.Println("services")
	var subTotal float64
	for _, value := range allS {
		if value.Minutes < 30 {
			subTotal += value.Price * float64(30)
		} else {
			subTotal += value.Price * float64(value.Minutes)
		}
	}
	cs <- subTotal
}

func totalMaintenance(allM []Maintenance, cm chan float64) {
	fmt.Println("maintenance")
	var subTotal float64
	for _, value := range allM {
		subTotal += value.Price
	}
	cm <- subTotal
}

func createProducts(cant int) []Products {
	var products = make([]Products, cant)
	for key, _ := range products {
		products[key] = Products{Name: fmt.Sprintf("ProductNumber_%d", key), Price: float64(rand.Intn(100-1) + 1), Cant: rand.Intn(100-1) + 1}
	}
	return products
}

func createServices(cant int) []Services {
	var services = make([]Services, cant)
	for key, _ := range services {
		services[key] = Services{Name: fmt.Sprintf("ServicesNumber_%d", key), Price: float64(rand.Intn(100-1) + 1), Minutes: rand.Intn(1000-1) + 1}
	}
	return services
}

func createMaintenances(cant int) []Maintenance {
	var maintenance = make([]Maintenance, cant)
	for key, _ := range maintenance {
		maintenance[key] = Maintenance{Name: fmt.Sprintf("MaintenanceNumber_%d", key), Price: float64(rand.Intn(100-1) + 1)}
	}
	return maintenance
}
