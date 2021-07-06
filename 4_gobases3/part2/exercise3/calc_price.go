package main

import (
	f "fmt"
)

func calcPrice(products []Product, services []Service, maintenances []Maintenance, ch chan float32) {
	go sumProducts(products, ch)
	go sumServices(services, ch)
	go sumMaintenance(maintenances, ch)
}

type Product struct {
	Name     string
	Price    float32
	Quantity int
}

type Service struct {
	Name        string
	Price       float32
	workMinutes int
}

type Maintenance struct {
	Name  string
	Price float32
}

func sumProducts(arrayProducts []Product, ch chan float32) {
	var totalPrice float32
	for _, product := range arrayProducts {
		totalPrice += product.Price * float32(product.Quantity)
	}
	ch <- totalPrice
}

func sumServices(arrayServices []Service, ch chan float32) {
	var totalServices float32
	for _, service := range arrayServices {
		totalServices += service.Price * (float32(service.workMinutes) / 30)
	}
	ch <- totalServices
}

func sumMaintenance(arrayMaintenance []Maintenance, ch chan float32) {
	var totalMaintenance float32
	for _, maintenance := range arrayMaintenance {
		totalMaintenance += maintenance.Price
	}
	ch <- totalMaintenance
}

func main() {
	ch := make(chan float32)
	var total float32

	products := []Product{
		{Name: "Soap", Price: 10, Quantity: 1},
		{Name: "Shampoo", Price: 20, Quantity: 2},
		{Name: "Coffe", Price: 30, Quantity: 5}}

	services := []Service{
		{Name: "Clean", Price: 20, workMinutes: 30},
		{Name: "Technical Service", Price: 50, workMinutes: 30}}
	maintenances := []Maintenance{
		{Name: "Fridge", Price: 70.5},
		{Name: "Washing machine", Price: 30.5}}

	calcPrice(products, services, maintenances, ch)

	for i := 0; i < 3; i++ {
		total += <-ch
	}
	f.Printf("La suma total es: %.2f\n", total)
}
