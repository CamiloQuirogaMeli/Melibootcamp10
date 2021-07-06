package main

import "fmt"

var total float64

func main() {

	products := []Product{{"Monitor", 500, 12}, {"Keyboard", 250, 21}}
	services := []Service{{"Technical service", 20, 30}, {"Computer assembly", 30, 60}}
	maintenances := []Maintenance{{"Preventive M¡maintenance", 30}, {"Computer corrective", 20}}

	c := make(chan float64)

	go SumProducts(products, c)
	go SumServices(services, c)
	go SumMaintenances(maintenances, c)

	for i := 0; i < 3; i++ {
		total += <-c
	}

	fmt.Println("The total sum is:", total)
}
