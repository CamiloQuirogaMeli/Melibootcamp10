package main

import (
	"fmt"
	"math"
)

func main() {
	products := []ProductA{{Name: "Mouse", Price: 250.00, Quantity: 3}, {Name: "Keyboard", Price: 1000.00, Quantity: 1}, {Name: "Trackpad", Price: 2000.00, Quantity: 1}}
	services := []Service{{Name: "Internet", Price: 3000.00, WorkedMinutes: 122}, {Name: "Water", Price: 400.00, WorkedMinutes: 62}}
	maintenances := []Maintenance{{Name: "House", Price: 2500.00}, {Name: "Car", Price: 10000.00}}

	p := make(chan float64)
	s := make(chan float64)
	m := make(chan float64)

	go sumPriceProducts(products, p)
	go sumPriceServices(services, s)
	go sumMaintenances(maintenances, m)

	total := <-p + <-s + <-m

	fmt.Printf("Total: [$%2.f]\n", total)
}

type ProductA struct {
	Name     string
	Price    float64
	Quantity int
}

type Service struct {
	Name          string
	Price         float64
	WorkedMinutes int
}

type Maintenance struct {
	Name  string
	Price float64
}

func sumPriceProducts(products []ProductA, c chan float64) {
	finalPrice := 0.0

	if len(products) == 0 {
		c <- finalPrice
	}

	for _, p := range products {
		finalPrice += p.Price * float64(p.Quantity)
	}

	fmt.Println("Products $", finalPrice)
	c <- finalPrice
}

func sumPriceServices(services []Service, c chan float64) {
	finalPrice := 0.0

	if len(services) == 0 {
		c <- finalPrice
	}

	for _, s := range services {
		wholeNumber, decimal := math.Modf(float64(s.WorkedMinutes) / 30)

		if decimal > 0.00 {
			wholeNumber++
		}

		finalPrice += s.Price * wholeNumber
	}

	fmt.Println("Services $", finalPrice)

	c <- finalPrice
}

func sumMaintenances(maintenances []Maintenance, c chan float64) {
	finalPrice := 0.0

	if len(maintenances) == 0 {
		c <- finalPrice
	}

	for _, m := range maintenances {
		finalPrice += m.Price
	}

	fmt.Println("Maintenances $", finalPrice)

	c <- finalPrice
}
