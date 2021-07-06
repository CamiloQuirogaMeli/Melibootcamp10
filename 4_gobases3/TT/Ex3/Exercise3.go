package main

import "fmt"

type product struct {
	name     string
	price    float64
	quantity int
}

type service struct {
	name    string
	price   float64
	minutes int
}

type maintenance struct {
	name  string
	price float64
}

func sumProd(products []product, c chan float64) {
	sum := 0.0

	for _, product := range products {
		sum += product.price * float64(product.quantity)
	}
	fmt.Println("Prod price:", sum)

	c <- sum
}

func sumServ(services []service, c chan float64) {
	sum := 0.0

	for _, service := range services {

		var halfHours int

		if service.minutes > 0 {
			halfHours = (service.minutes / 30) + 1
		}

		sum += service.price * float64(halfHours)
	}
	fmt.Println("Serv price:", sum)

	c <- sum

}

func sumMaint(maintenances []maintenance, c chan float64) {

	sum := 0.0

	for _, maintenance := range maintenances {
		sum += maintenance.price
	}

	fmt.Println("Maint price:", sum)

	c <- sum

}

func main() {

	prod1 := product{"Prod 1", 100.00, 5}
	prod2 := product{"Prod 2", 20.00, 1}
	prod3 := product{"Prod 3", 1000.00, 10}

	prods := []product{prod1, prod2, prod3}

	serv1 := service{"serv 1", 10.00, 30}
	serv2 := service{"serv 2", 100.00, 100}
	serv3 := service{"serv 3", 50.00, 60}

	servs := []service{serv1, serv2, serv3}

	maint1 := maintenance{"Maint 1", 100.00}
	maint2 := maintenance{"Maint 2", 1.00}
	maint3 := maintenance{"Maint 3", 20.00}

	maints := []maintenance{maint1, maint2, maint3}

	chanProds := make(chan float64)
	chanServs := make(chan float64)
	chanMaints := make(chan float64)

	go sumProd(prods, chanProds)
	go sumServ(servs, chanServs)
	go sumMaint(maints, chanMaints)

	prodsPrice := <-chanProds
	servsPrice := <-chanServs
	maintsPrice := <-chanMaints

	fmt.Println("Total proce:", prodsPrice+servsPrice+maintsPrice)
}
