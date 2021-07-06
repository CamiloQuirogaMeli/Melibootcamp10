package main

import "fmt"

type Product struct {
	name   string
	price  float64
	amount int
}

type Service struct {
	name     string
	price    float64
	duration int
}

type Manteinance struct {
	name  string
	price float64
}

func main() {

	c := make(chan string)
	go func(c chan string) {

		products := hardCodedProducts()
		total := addProducts(products)
		c <- fmt.Sprint("Total products:", total)

	}(c)

	go func(c chan string) {

		services := hardCodedServices()
		total := addServices(services)
		c <- fmt.Sprint("Total services:", total)

	}(c)

	go func(c chan string) {

		mants := hardCodedMant()
		total := addManteinance(mants)
		c <- fmt.Sprint("Total mants:", total)

	}(c)

	syncronizeString(3, c)

}

func addProducts(products []Product) float64 {

	total := 0.0
	c := make(chan float64)
	for _, product := range products {
		go func(product Product, c chan float64) {
			total += product.price * float64(product.amount)
			c <- product.price * float64(product.amount)
		}(product, c)
	}

	syncronize(len(products), c)
	return total
}

func addServices(services []Service) float64 {

	total := 0.0
	c := make(chan float64)
	for _, service := range services {
		go func(service Service, c chan float64) {
			duration := service.duration
			halfHours := 1.0
			if duration > 30 {
				halfHours = float64(duration) / 30
			}
			total += service.price * halfHours
			c <- service.price * halfHours
		}(service, c)
	}
	syncronize(len(services), c)
	return total
}

func addManteinance(mants []Manteinance) float64 {

	total := 0.0
	c := make(chan float64)
	for _, mant := range mants {
		go func(mant Manteinance, c chan float64) {
			total += mant.price
			c <- mant.price
		}(mant, c)
	}

	syncronize(len(mants), c)
	return total
}

func syncronize(numGoroutines int, c chan float64) {
	for i := 0; i < numGoroutines; i++ {
		//fmt.Println("function just recieved this: ", <-c)
		<-c
	}
}

func syncronizeString(numGoroutines int, c chan string) {
	for i := 0; i < numGoroutines; i++ {
		fmt.Println("function just recieved this: ", <-c)
	}
}

func hardCodedProducts() []Product {
	p1 := Product{"a", 200, 2}
	p2 := Product{"b", 100, 1}
	return []Product{p1, p2}
}

func hardCodedServices() []Service {
	s1 := Service{"a", 100, 60}
	s2 := Service{"b", 100, 30}
	return []Service{s1, s2}
}
func hardCodedMant() []Manteinance {
	m1 := Manteinance{"a", 100}
	m2 := Manteinance{"b", 100}
	return []Manteinance{m1, m2}
}
