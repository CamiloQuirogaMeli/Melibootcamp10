package ecommerce

type StoreSponge struct {
	size string
	shippingAddress string
}

func (store StoreSponge) Price() float64 {
	const PRICE = 1000
	switch store.size {
	case SMALL_PRODUCT:
		return PRICE
	case MEDIUM_PRODUCT:
		return PRICE + PRICE * 0.03
	case LARGE_PRODUCT:
		return PRICE + PRICE * 0.06 + 2500.0
	default:
		return -1
	}
}

func (store StoreSponge) ShippingAddress() string {
	return store.shippingAddress
}