package main

import (
	"fmt"
	"math"
)

type Product struct {
	id     string
	price  float64
	amount int
}

func sumProducts(products []Product, c chan float64) {
	var total float64 = 0
	for _, product := range products {
		total += product.price * float64(product.amount)
	}
	c <- total
}

type Service struct {
	id      string
	price   float64
	minutes int
}

func sumService(services []Service, c chan float64) {
	var total float64 = 0
	for _, service := range services {
		workTime := math.Ceil(float64(service.minutes) / 30)
		total += service.price * workTime
	}

	c <- total
}

type Maintenance struct {
	id    string
	price float64
}

func sumMaintenances(maintenances []Maintenance, c chan float64) {
	var total float64 = 0
	for _, maintenance := range maintenances {

		total += maintenance.price
	}
	c <- total
}

func main() {

	var option int
	var aux int = 1
	var products []Product

	products = append(products, Product{"1", 10000.0, 2})
	products = append(products, Product{"2", 20000.0, 3})
	products = append(products, Product{"3", 30000.0, 4})

	var services []Service

	services = append(services, Service{"1", 10000.0, 3})
	services = append(services, Service{"2", 20000.0, 3})
	services = append(services, Service{"3", 30000.0, 3})

	var maintenances []Maintenance
	maintenances = append(maintenances, Maintenance{"1", 10000.0})
	maintenances = append(maintenances, Maintenance{"2", 20000.0})
	maintenances = append(maintenances, Maintenance{"3", 30000.0})
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Crear producto\n2. Crear servicio\n3. Crear mantenimiento\n4. Calcular Sumatoria \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			product := createProduct()
			products = append(products, product)

			fmt.Println("El producto ha sido creado con exito")
		case 2:
			service := createService()
			services = append(services, service)

			fmt.Println("El servicio ha sido creado con exito")
		case 3:
			maintenance := createMaintenance()
			maintenances = append(maintenances, maintenance)

			fmt.Println("El mantenimiento ha sido creado con exito")
		case 4:
			calculateSum(products, services, maintenances)
		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}

}

func calculateSum(products []Product, services []Service, maintenances []Maintenance) {
	p := make(chan float64)
	s := make(chan float64)
	m := make(chan float64)

	go sumProducts(products, p)
	go sumService(services, s)
	go sumMaintenances(maintenances, m)

	var total float64
	total += <-p + <-s + <-m

	fmt.Printf("Total: $%.2f", total)
}

func createProduct() Product {
	var id string
	var price float64
	var amount int
	fmt.Println("Ingrese el Id del producto")
	fmt.Scan(&id)
	fmt.Println("Ingrese el precio del producto")
	fmt.Scan(&price)
	fmt.Println("Ingrese la cantidad de productos")
	fmt.Scan(&amount)

	return Product{id, price, amount}

}

func createService() Service {
	var id string
	var price float64
	var minutes int
	fmt.Println("Ingrese el Id del producto")
	fmt.Scan(&id)
	fmt.Println("Ingrese el precio del producto")
	fmt.Scan(&price)
	fmt.Println("Ingrese la cantidad de minutos trabajados")
	fmt.Scan(&minutes)

	return Service{id, price, minutes}

}

func createMaintenance() Maintenance {
	var id string
	var price float64
	fmt.Println("Ingrese el Id del producto")
	fmt.Scan(&id)
	fmt.Println("Ingrese el precio del producto")
	fmt.Scan(&price)

	return Maintenance{id, price}

}
