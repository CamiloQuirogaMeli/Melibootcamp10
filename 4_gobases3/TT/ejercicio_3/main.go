package main

import "fmt"


func main() {

	products := []Product{{"botella", 15.5, 5}, {"lechuga", 80, 10}}
	services := []Service{{"limpieza", 20.5, 100}, {"servicio tecnico", 50, 120}}
	maintenances := []Maintenance{{"horno", 70.5}, {"microondas", 30.5}}

	ch := make(chan float64)
	go SumProducts(products, ch)
	go SumServices(services, ch)
	go SumMaintenances(maintenances, ch)

	var total float64
	for i := 0; i < 3; i++ {
		total += <- ch
	}

	fmt.Println("El monto final es:", total)
}