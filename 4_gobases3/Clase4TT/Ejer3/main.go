package main

import "fmt"

type product struct {
	name   string
	price  float64
	amount int64
}
type service struct {
	name    string
	price   float64
	minutes int64
}

type maintenance struct {
	name  string
	price float64
}

func sumProducts(products []product, c chan float64) {
	var sum float64 = 0
	for _, item := range products {
		sum += item.price * float64(item.amount)
	}
	c <- sum
}

func sumServices(service []service, c chan float64) {
	var sum float64 = 0
	for _, item := range service {
		if item.minutes < 30 {
			sum += item.price
		} else {
			sum += item.price * float64(item.minutes) / 30
		}
	}
	c <- sum
}

func sumMaintenance(maintenances []maintenance, c chan float64) {
	var sum float64 = 0
	for _, item := range maintenances {
		sum += item.price
	}
	c <- sum
}

func main() {
	c1, c2, c3 := make(chan float64), make(chan float64), make(chan float64)
	go sumProducts([]product{{"p1", 1, 1}, {"p2", 2, 2}, {"p3", 3, 3}}, c1)
	go sumServices([]service{{"s1", 1000, 10}, {"s2", 2, 20}, {"s3", 3, 30}}, c2)
	go sumMaintenance([]maintenance{{"m1", 1}, {"m2", 2}, {"m3", 3}}, c3)
	total := <-c1 + <-c2 + <-c3
	fmt.Println(total)
}
