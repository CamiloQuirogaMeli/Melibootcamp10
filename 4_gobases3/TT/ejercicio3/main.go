package main

import "fmt"

func main() {

	products := []Product{
		{"p1", 10.0, 1},
	}
	services := []Service{
		{"s1", 15, 65},
	}
	maintenances := []Maintenance{
		{price: 50, name: "m1"},
	}

	channel := make(chan float64)
	go sumServices(services, channel)
	go sumMaintenances(maintenances, channel)
	go sumProducts(products, channel)

	var total float64
	for i := 0; i < 3; i++ {
		total += <- channel
	}

	fmt.Println("el precio final es:", total)
}
