package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	small  = "small"
	medium = "medium"
	big    = "big"
	//smallStore: It only has one product, and said product can be only small. Also, it doesn't make shipments
	smallStore = "small"
	//bigStore: It only has one product, and said product can be only big.
	bigStore = "big"
)

//********** PRODUCT ****************//
type Product struct {
	baseCost    float64
	productType string
	cost        float64
}

func (product Product) toString() string {
	var detail string = ""
	detail += "Base cost " + strconv.FormatFloat(product.baseCost, 'f', -1, 64) + "\n"
	detail += "Type of product " + product.productType + "\n"
	detail += "Final cost:" + strconv.FormatFloat(product.cost, 'f', -1, 64) + "\n"

	return detail

}
func newProduct(theBaseCost float64, theProductType string) (Product, error) {

	product := Product{}

	aditionalPercentageCost := map[string]float64{small: 0, medium: 0.03, big: 0.06}
	aditionalShippingCost := map[string]float64{small: 0, medium: 0, big: 2500}

	if _, ok := aditionalPercentageCost[theProductType]; !ok {
		return product, errors.New(theProductType + " is not a valid product type")
	}

	product.baseCost = theBaseCost
	product.productType = theProductType
	product.cost = theBaseCost*(1+aditionalPercentageCost[theProductType]) + aditionalShippingCost[theProductType]

	return product, nil
}

//********** PRODUCT ****************//
//********** STORE ****************//
type Ecommerce interface {
	getCost() float64
	getShippingAddress() string
	addProduct(product Product) error
	toString() string
}

type TiendaDeBarrio struct {
	name    string
	product Product
}

func (store TiendaDeBarrio) getCost() float64 {
	return store.product.cost
}
func (store TiendaDeBarrio) getShippingAddress() string {
	return store.name + " doesn't do shipments"
}

func (store *TiendaDeBarrio) addProduct(product Product) error {
	if product.productType == small {
		store.product = product
		return nil
	}
	return errors.New(store.name + " doesn't sell " + product.productType + " products")
}

func (store TiendaDeBarrio) toString() string {

	var detail string = ""
	detail += "MiniMarket: " + store.name + "\n"
	detail += "Shipping address " + store.getShippingAddress() + "\n"
	detail += "Product info:\n"
	detail += store.product.toString()

	return detail
}

type Supermercado struct {
	name            string
	shippingAddress string
	product         Product
}

func (store Supermercado) getCost() float64 {
	return store.product.cost
}
func (store Supermercado) getShippingAddress() string {
	return store.shippingAddress
}
func (store *Supermercado) addProduct(product Product) error {

	if product.productType == big {
		store.product = product
		return nil
	}

	return errors.New(store.name + " doesn't sell " + product.productType + " products")
}

func (store Supermercado) toString() string {

	var detail string = ""
	detail += "Supermarket " + store.name + "\n"
	detail += "Shipping address " + store.getShippingAddress() + "\n"
	detail += "Product info:\n"
	detail += store.product.toString()

	return detail
}

func storeFactory(storeType string, params ...string) (Ecommerce, error) {

	switch storeType {
	case smallStore:
		return &TiendaDeBarrio{name: params[0]}, nil
	case bigStore:
		return &Supermercado{name: params[0], shippingAddress: params[1]}, nil
	default:
		return &TiendaDeBarrio{}, errors.New("store Type " + storeType + " is not valid")
	}

}

//********** STORE ****************//

func main() {

	carulla, err := storeFactory("big", "Carulla", "Cra 32 no 53 23")
	if err != nil {
		fmt.Println(err)
	}
	productForCarulla, err := newProduct(1000, "big")
	if err != nil {
		fmt.Println(err)
	} else {
		err = carulla.addProduct(productForCarulla)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(carulla.toString())
		}

	}

	miCasita, err := storeFactory("small", "Mi casita", "Cra 12 no 54 83")
	if err != nil {
		fmt.Println(err)
	}
	productFormiCasita, err := newProduct(1000, "small")
	if err != nil {
		fmt.Println(err)
	} else {
		err = miCasita.addProduct(productFormiCasita)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(miCasita.toString())
		}

	}

}
