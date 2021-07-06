package main

import "fmt"

type Product struct {
	name  string
	price float64
	total int
}

type Service struct {
	name    string
	price   float64
	minWork int
}

type Maintenance struct {
	name  string
	price float64
}

func SumProducts(p []Product, c chan float64) {
	var sum float64
	for _, value := range p {
		price := value.price * float64(value.total)
		sum += price
	}
	c <- sum
}

func SumServices(s []Service, c chan float64) {
	var sum float64
	for _, value := range s {
		price := value.price * float64(value.minWork)
		sum += price
	}
	c <- sum
}

func sumMaintenance(m []Maintenance, c chan float64) {
	var sum float64
	for _, valor := range m {
		sum += valor.price
	}
	c <- sum
}

func main() {
	productos := []Product{
		{"p1", 123.45, 1},
		{"p2", 234.56, 2},
		{"p3", 45.67, 3},
		{"p4", 78.90, 4}}
	servicios := []Service{
		{"s1", 100, 90},
		{"s2", 200, 45},
		{"s3", 300, 60}}
	manteniminetos := []Maintenance{
		{"m1", 2000},
		{"m2", 1500}}

	channel1 := make(chan float64)
	channel2 := make(chan float64)
	channel3 := make(chan float64)

	go SumProducts(productos, channel1)
	go SumServices(servicios, channel2)
	go sumMaintenance(manteniminetos, channel3)

	ch1 := <-channel1
	ch2 := <-channel2
	ch3 := <-channel3

	fmt.Println("El precio total de los productos:", ch1)
	fmt.Println("El precio total de los servicios:", ch2)
	fmt.Println("El precio total de los mantenimientos:", ch3)

	fmt.Println("La suma de todos los precios es: ", ch1+ch2+ch3)
}
