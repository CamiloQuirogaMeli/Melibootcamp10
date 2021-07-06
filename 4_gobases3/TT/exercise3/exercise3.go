package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

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

func sumProducts(products []Product, c chan float64) {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.price * float64(product.quantity)
	}
	// defer wg.Done()
	c <- totalPrice
}

func sumServices(services []Service, c chan float64) {
	totalPrice := 0.0
	for _, service := range services {
		totalPrice += service.price * float64(service.minutesWorked)
	}
	// defer wg.Done()
	c <- totalPrice
}

func sumMaintenances(maintenances []Maintenance, c chan float64) {
	totalPrice := 0.0
	for _, maintenance := range maintenances {
		totalPrice += maintenance.price
	}
	// defer wg.Done()
	c <- totalPrice
}

func main() {

	products := []Product{{"Hammer", 123.3, 5}, {"Camera", 234.43, 5}}
	services := []Service{{"Landing page creation", 200, 30}, {"Shipping service", 20, 30}}
	maintenances := []Maintenance{{"Air conditioner", 2000}, {"Mow the lawn", 1500}}

	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)
	go sumProducts(products, c1)
	go sumServices(services, c2)
	go sumMaintenances(maintenances, c3)

	totalSum := <-c1 + <-c2 + <-c3
	fmt.Println(totalSum)

	// c := make(chan interface{}, 3)
	// wg.Add(3)
	// go sumProducts(products, c)
	// go sumServices(services, c)
	// go sumMaintenances(maintenances, c)
	// wg.Wait()
	// close(c)
	// for item := range c {
	// 	fmt.Println(item)
	// }

}
