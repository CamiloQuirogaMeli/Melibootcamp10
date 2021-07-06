package product

import "errors"

type Product interface {
	TotalSize() int
	IsSpecial() bool
}

const (
	SMALL_PRODUCT 	= "Small Product"
	MEDIUM_PRODUCT 	= "Medium Product"
	LARGE_PRODUCT 	= "Large Product"
	FRAGILE_PRODUCT = "Fragile Product"
	SPECIAL_PRODUCT = "Special Product"
)

func ProductFactory(productType string, size int) (product Product, err error) {
	switch productType {
		case SMALL_PRODUCT:
			product = SmallProduct{size}
		case MEDIUM_PRODUCT:
			product = MediumProduct{size}
		case LARGE_PRODUCT:
			product = LargeProduct{size}
		case FRAGILE_PRODUCT:
			product = FragileProduct{size}
		case SPECIAL_PRODUCT:
			product = SpecialProduct{size}
		default:
			err = errors.New("El tipo de producto es invalido")
	}
	return
}