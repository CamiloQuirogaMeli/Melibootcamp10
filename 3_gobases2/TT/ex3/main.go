package main

import (
	"fmt"
)

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

type Ecommerce interface {
	getPrice()
	getAddress()
}

type StoreOne struct {
	Name    string
	Address string
	Price   float64
}

type StoreTwo struct {
	Name    string
	Address string
	Price   float64
}

func (s StoreOne) getPrice() {
	fmt.Println(s.Price)
}

func (s StoreOne) getAddress() {
	fmt.Println(s.Address)
}

func (s StoreTwo) getPrice() {
	fmt.Println(s.Price)
}

func (s StoreTwo) getAddress() {
	fmt.Println(s.Address)
}

func getPrice(price float64, size string) float64 {

	switch size {
	case SMALL:
		return price
	case MEDIUM:
		return price * 1.03
	case LARGE:
		return (price * 1.06) + 2500
	}

	return 0
}

func newStore(name string, storeType int, price float64, size string, address string) Ecommerce {

	total := getPrice(price, size)

	switch storeType {
	case 1:
		return StoreOne{name, address, total}
	case 2:
		return StoreTwo{name, address, total}
	}

	return nil
}

func main() {
	var name string
	var storeType int
	var price float64
	var size int
	var address string

	fmt.Print("Esciba el nombre de la tienda: ")
	fmt.Scanln(&name)
	fmt.Print("Esciba el tipo de la tienda (1 o 2): ")
	fmt.Scanln(&storeType)
	fmt.Print("Esciba el precio: ")
	fmt.Scanln(&price)
	fmt.Print("Esciba el numero que represente el tamaño (1: pequeño, 2: mediano, 3: grande): ")
	fmt.Scanln(&size)
	fmt.Print("Esciba la dirección: ")
	fmt.Scanln(&address)

	switch size {
	case 1:
		store := newStore(name, 1, price, SMALL, address)
		store.getPrice()
		store.getAddress()
	case 2:
		store := newStore(name, 1, price, MEDIUM, address)
		store.getPrice()
		store.getAddress()
	case 3:
		store := newStore(name, 1, price, LARGE, address)
		store.getPrice()
		store.getAddress()
	}
}
