package main

import (
	"fmt"
)

type Product struct {
	name     string
	price    float64
	quantity int
}

type Service struct {
	name          string
	price         float64
	minutesWorked int
}

type Maintenance struct {
	name  string
	price float64
}

func addProducts(products []Product, c chan float64) {

	var totalPrice float64
	for _, product := range products {
		totalPrice += product.price
	}

	c <- totalPrice
}

func addServices(services []Service, c chan float64) {

	var totalPrice float64
	for _, service := range services {

		halfHourWorked := service.minutesWorked / 30
		totalPrice += service.price * float64(halfHourWorked)
	}

	c <- totalPrice
}

func addMaintenance(maintenances []Maintenance, c chan float64) {

	var totalPrice float64
	for _, maintenance := range maintenances {
		totalPrice += maintenance.price
	}

	c <- totalPrice
}
func main() {

	var totalPrice float64
	channel := make(chan float64)

	product1 := Product{"doritos", 200, 3}
	product2 := Product{"monster", 150, 6}
	products := []Product{product1, product2}

	service1 := Service{"netflix", 900, 120}
	service2 := Service{"disney", 360, 360}
	services := []Service{service1, service2}

	maintenance1 := Maintenance{"plomero", 299}
	maintances := []Maintenance{maintenance1}

	go addProducts(products, channel)
	totalPrice += <-channel
	go addServices(services, channel)
	totalPrice += <-channel
	go addMaintenance(maintances, channel)
	totalPrice += <-channel
	fmt.Print(totalPrice)
}
