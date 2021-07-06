package ecommerce

import "errors"

const (
	SMALL_PRODUCT 	= "Small Product"
	MEDIUM_PRODUCT = "Medium Product"
	LARGE_PRODUCT = "Large Product"
)

const (
	SPONGE = "Sponge"
	STAR = "Star"
)

type Ecommerce interface {
	Price() float64
	ShippingAddress() string
}

func NewStore(size string, storeType string, shippingAddress string) (Ecommerce, error) {
	switch storeType {
		case SPONGE:
			return StoreSponge{size, shippingAddress}, nil
		case STAR:
			return StoreStar{size, shippingAddress}, nil
		default:
			return nil, errors.New("El tipo de tienda no existe")
	}
}