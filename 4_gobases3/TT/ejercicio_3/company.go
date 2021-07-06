package main

import "math"

type Product struct {
	Name		string
	Price		float64
	Quantity	int
}

type Service struct {
	Name			string
	Price			float64
	MinutesWorked	int
}

type Maintenance struct {
	Name	string
	Price	float64
}

func SumProducts(products []Product, ch chan float64) {
	var total float64
	for _, product := range products {
		total += product.Price * float64(product.Quantity)
	}
	ch <- total
}

func SumServices(services []Service, ch chan float64) {
	var total float64
	for _, service := range services {
		total += service.Price * math.Ceil(float64(service.MinutesWorked) / 30.0)
	}
	ch <- total
}

func SumMaintenances(maintenances []Maintenance, ch chan float64) {
	var total float64
	for _, maintenance := range maintenances {
		total += maintenance.Price
	}
	ch <- total
}