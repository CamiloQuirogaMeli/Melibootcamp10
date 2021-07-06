package main

import (
	"fmt"
)

const (
	SMALL  = "pequeño"
	MEDIUM = "mediano"
	BIG    = "grande"
)

type AbstractShop struct {
	Address string
	Phone   string
	Website string
}

func (as AbstractShop) price(productType string, productPrice float64) float64 {
	switch productType {
	case SMALL:
		return productPrice
	case MEDIUM:
		return productPrice * 1.03
	case BIG:
		return productPrice*1.03 + 2500
	default:
		return 0
	}
}

func (as AbstractShop) delivery(deliveryAddress string) string {
	return deliveryAddress
}

type ShopA struct {
	AbstractShop
	ClosingTime int
}

type ShopB struct {
	AbstractShop
	DeliveryMaxDistanceKm int
}

type Ecommerce interface {
	price(string, float64) float64
	delivery(string) string
}

func newShop(shopType string) Ecommerce {
	switch shopType {
	case "shopA":
		return ShopA{AbstractShop{"cll 83 12 14", "12345", "shopA.com"}, 18}
	case "shopB":
		return ShopB{AbstractShop{"cll 33 12 14", "12345", "shopB.com"}, 9}
	default:
		return nil
	}
}

func main() {
	shopType := "shopA"
	//shopType := "shopB"
	var shop Ecommerce = newShop(shopType)
	fmt.Println("El precio del producto en la tinda", shopType, "es : ", shop.price(SMALL, 2000))
	fmt.Println("La dirección de entrga es", shop.delivery("cll 89 33 45"))

}
