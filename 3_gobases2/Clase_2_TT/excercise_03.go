package main

import (
	"fmt"
)

func main() {
	store := newStore("StoreOne")
	cost := store.Price()
	shipment := store.Shipment()

	fmt.Printf("The final cost: [%2.f]\n", cost)
	fmt.Printf("Address Shipment: [%s]\n", shipment)
}

type Ecommerce interface {
	Price() float64
	Shipment() string
}

type StoreOne struct {
	TypeProduct string
	Cost        float64
	Address     string
}

type StoreTwo struct {
	TypeProduct string
	Cost        float64
	Address     string
}

func newStore(store string) Ecommerce {
	if store == "StoreOne" {
		return StoreOne{TypeProduct: "large", Cost: 5000.00, Address: "Chinatown2"}
	}
	if store == "StoreTwo" {
		return StoreTwo{TypeProduct: "medium", Cost: 2500.00, Address: "San Francisco"}
	}
	return nil
}

func (p StoreOne) Price() float64 {
	switch p.TypeProduct {
	case "small":
		return p.Cost
	case "medium":
		aditional := p.Cost * 0.03
		return p.Cost + aditional
	case "large":
		aditional := p.Cost*0.06 + 2500.00
		return p.Cost + aditional
	}
	return 0.0
}

func (p StoreOne) Shipment() string {
	return p.Address
}

func (p StoreTwo) Price() float64 {
	switch p.TypeProduct {
	case "small":
		return p.Cost
	case "medium":
		aditional := p.Cost * 0.03
		return p.Cost + aditional
	case "large":
		aditional := p.Cost*0.06 + 2500.00
		return p.Cost + aditional
	}
	return 0.0
}

func (p StoreTwo) Shipment() string {
	return p.Address
}
