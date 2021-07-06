package main

import "fmt"

type Product struct {
	price float64
	typeProduct string
}

type Stores struct {
	Product
	typeStore string
	address string
}

type StoreOne struct {
	Stores
}

type StoreTwo struct {
	Stores
}

type Ecommerce interface {
	getPrice() float64
	getShipping() string
}

const (
	SMALL     = "small"
	MEDIUM    = "medium"
	LARGE     = "large"
	STORE_ONE = "storeOne"
	STORE_TWO = "storeTwo"
)
func main(){
	store := newStore(STORE_TWO)
	fmt.Println(store.getShipping())
	fmt.Println(store.getPrice())
}

func newStore(typeStore string) Ecommerce{
	var store Ecommerce
	switch typeStore {
	case STORE_ONE:
		store = StoreOne{
			Stores{
				Product{price: 15.11, typeProduct: LARGE},
				STORE_ONE,
				"Address 1",
			},
		}
	case STORE_TWO:
		store = StoreTwo{
			Stores{
				Product{price:8.4, typeProduct: SMALL},
				STORE_TWO,
				"Address 2",
			},
		}
	}
	return store
}

func (store Stores) getPrice() float64 {
	switch store.Product.typeProduct {
	case MEDIUM:
		return store.Product.price * 1.03
	case LARGE:
		return store.Product.price * 1.06 + 2500.0
	default:
		return store.Product.price
	}
}

func (store Stores) getShipping() string {
	return store.address
}