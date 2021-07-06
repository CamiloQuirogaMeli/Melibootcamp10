package main

import "fmt"

type Product struct {
	name string
	price float64
	amount int
}
type Service struct {
	name string
	price float64
	minutesWorked int
}
type Maintenance struct {
	name string
	price float64
}

func sumProducts(products []Product, channel chan float64){
	var total float64
	for _, product := range products {
		total += product.price * float64(product.amount)
	}
	channel <- total
}

func sumServices(services []Service, channel chan float64) {
	var total float64
	for _, service := range services {
		total += service.price * float64(int(float64(service.minutesWorked) / 30.0))
	}
	fmt.Println(total)
	channel <- total
}

func sumMaintenances(maintenances []Maintenance, channel chan float64) {
	var total float64
	for _, maintenance := range maintenances {
		total += maintenance.price
	}
	channel <- total
}