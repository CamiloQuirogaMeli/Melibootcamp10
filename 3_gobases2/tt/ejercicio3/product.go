package ejercicio3

import (
	"errors"
	"fmt"
	"strings"
)

const (
	SMALL  = "small"
	MEDIUM = "medium"
	BIG    = "big"
)

type (
	Product struct {
		Kind      string
		UnitPrice float64
	}
	Company struct {
		Name    string
		Address string
	}
	Ecommerce interface {
		Price(product Product) (float64, error)
		Shipping(address string) string
	}
)

func (c Company) Price(p Product) (total float64, err error) {
	p.Kind = strings.ToUpper(p.Kind)
	switch p.Kind {
	case SMALL:
		total = p.UnitPrice
	case MEDIUM:
		total = p.UnitPrice + (p.UnitPrice * 0.03)
	case BIG:
		total = p.UnitPrice + (p.UnitPrice * 0.06) + 2500
	default:
		err = errors.New("product kind not exists")
	}
	return
}

func (c Company) Shipping(address string) string {
	return address
}

func NewShop(name string, storeType int, address string) Ecommerce {
	switch storeType {
	case 1:
		return Company{name, address}
	case 2:
		return Company{name, address}
	}
	return nil
}

func main() {
	shopOne := NewShop("ShopOne", 1, "Cra 9")
	shopTwo := NewShop("ShopTwo", 2, "Cra 9")

	productBig := Product{
		Kind:      BIG,
		UnitPrice: 9000,
	}
	productMedium := Product{
		Kind:      MEDIUM,
		UnitPrice: 4000,
	}
	productSmall := Product{
		Kind:      SMALL,
		UnitPrice: 2000,
	}

	fmt.Println(shopOne.Price(productSmall))
	fmt.Println(shopOne.Price(productMedium))
	fmt.Println(shopOne.Price(productBig))

	fmt.Println(shopTwo.Price(productSmall))
	fmt.Println(shopTwo.Price(productMedium))
	fmt.Println(shopTwo.Price(productBig))

	fmt.Println(shopOne.Shipping("Calle 100"))
	fmt.Println(shopTwo.Shipping("Calle 134"))
}
