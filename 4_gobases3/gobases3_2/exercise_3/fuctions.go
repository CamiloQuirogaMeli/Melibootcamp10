package main

type Product struct {
	name   string
	price  float64
	amount int
}

type Service struct {
	name          string
	price         float64
	minutesWorked float64
}

type Maintenance struct {
	name  string
	price float64
}

func SumProducts(products []Product, c chan float64) {

	var total float64
	for _, product := range products {
		total += product.price * float64(product.amount)
	}

	c <- total

}

func SumServices(services []Service, c chan float64) {

	var total float64
	for _, service := range services {
		total += service.price * service.minutesWorked / 30.0
	}

	c <- total

}

func SumMaintenances(maintenances []Maintenance, c chan float64) {

	var total float64
	for _, maintenance := range maintenances {
		total += maintenance.price
	}

	c <- total

}
