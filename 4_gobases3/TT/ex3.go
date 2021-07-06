package main

import (
	"fmt"
	"math"
)

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type Service struct {
	Name          string
	Price         float64
	MinutesWorked int
}

type Maintenance struct {
	Name  string
	Price float64
}

func totalPriceProducts(products []Product, c chan float64) float64 {
	total := 0.0
	for _, product := range products {
		total += product.Price * float64(product.Amount)
	}
	fmt.Println("total de los productos", total)
	c <- total
	return total
}

func totalPriceServices(services []Service, c chan float64) float64 {
	total := 0.0
	for _, service := range services {
		total += service.Price * math.Ceil(float64(service.MinutesWorked)/30.0)
	}
	fmt.Println("total de los servicios", total)
	c <- total
	return total
}

func totalPriceMaintenances(maintenances []Maintenance, c chan float64) float64 {
	total := 0.0
	for _, maintenance := range maintenances {
		total += maintenance.Price
	}
	fmt.Println("total de los mantenimientos", total)
	c <- total
	return total
}

func main() {
	c := make(chan float64)
	maintenances := []Maintenance{{"a", 1}, {"a", 1}, {"a", 1}, {"a", 1}, {"a", 1}, {"a", 1}}
	products := []Product{{"x", 1, 2}, {"x", 1, 2}, {"x", 1, 2}, {"x", 1, 2}, {"x", 1, 2}, {"x", 1, 2}, {"x", 1, 2}}
	services := []Service{{"y", 2, 45}, {"y", 2, 45}, {"y", 2, 45}, {"y", 2, 45}, {"y", 2, 45}, {"y", 2, 45}}
	go totalPriceMaintenances(maintenances, c)
	go totalPriceProducts(products, c)
	go totalPriceServices(services, c)
	totalPrice := 0.0
	totalPrice += <-c
	totalPrice += <-c
	totalPrice += <-c
	fmt.Println("El precio total de todo es", totalPrice)
}
