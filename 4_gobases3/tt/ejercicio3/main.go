package main

import (
	"fmt"
	"sync"
)

type Product struct {
	name   string
	price  float64
	amount float64
}

type Service struct {
	name       string
	price      float64
	activeTime float64
}

type Maintenance struct {
	name  string
	price float64
}

var wg sync.WaitGroup // 1

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
func TotalPriceProducts(tch *float64, products [2]Product) {
	var t float64

	for _, p := range products {
		t += p.price * p.amount
	}
	*tch += t
	wg.Done()
}

func TotalPriceServices(tch *float64, services [2]Service) {

	var t float64
	for _, s := range services {
		workByMinutes := 30.0
		if s.activeTime > 30 {
			workByMinutes = s.activeTime
		}
		t += s.price * workByMinutes
	}
	*tch += t
	wg.Done()
}

func TotalPriceMaintenance(tch *float64, mt [2]Maintenance) {
	var t float64
	for _, m := range mt {
		t += m.price
	}

	*tch += t
	wg.Done()

}

func main() {
	services := generateSvc()
	maintenance := generateMt()
	products := generateProd()
	var totalPrice float64

	wg.Add(1)
	go TotalPriceProducts(&totalPrice, products)
	wg.Add(1)
	go TotalPriceMaintenance(&totalPrice, maintenance)
	wg.Add(1)
	go TotalPriceServices(&totalPrice, services)
	wg.Wait()

	fmt.Println(totalPrice)
}

func generateProd() [2]Product {
	var products [2]Product
	p1 := Product{
		name:   "fideos",
		price:  50.0,
		amount: 3,
	}
	p2 := Product{
		name:   "arroz",
		price:  70.0,
		amount: 3,
	}
	products[0] = p1
	products[1] = p2
	return products
}

func generateSvc() [2]Service {
	var service [2]Service
	p1 := Service{
		name:       "fideos",
		price:      50.0,
		activeTime: 60,
	}
	p2 := Service{
		name:       "arroz",
		price:      70.0,
		activeTime: 60,
	}
	service[0] = p1
	service[1] = p2
	return service
}

func generateMt() [2]Maintenance {
	var maintenance [2]Maintenance
	p1 := Maintenance{
		name:  "fideos",
		price: 50.0,
	}
	p2 := Maintenance{
		name:  "arroz",
		price: 70.0,
	}
	maintenance[0] = p1
	maintenance[1] = p2
	return maintenance
}
