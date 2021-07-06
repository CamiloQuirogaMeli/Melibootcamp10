package main

import "fmt"

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

type Ecommerce interface {
	storePrice() float64
	shipping() string
	addProduct() Product
}

type Product interface {
	price() float64
}

type Store1 struct {
	name     string
	color    string
	products []Product
}

type Store2 struct {
	name     string
	products []Product
}

type Small struct {
	name      string
	basePrice float64
}

type Medium struct {
	name      string
	basePrice float64
}

type Large struct {
	name      string
	basePrice float64
}

func (s1 *Store1) addProduct(productType string, price float64, productName string) Product {
	newProduct := calculateFinalPrice(productType, price, productName)
	s1.products = append(s1.products, newProduct)
	return newProduct
}

func (s2 *Store2) addProduct(productType string, price float64, productName string) Product {
	newProduct := calculateFinalPrice(productType, price, productName)
	s2.products = append(s2.products, newProduct)
	return newProduct
}

func (s1 Store1) storePrice() float64 {
	var total float64 = 0
	for _, product := range s1.products {
		total += float64(product.price())
	}
	return total
}

func (s2 Store2) storePrice() float64 {
	var total float64 = 0
	for _, product := range s2.products {
		total += float64(product.price())
	}
	return total
}

func (s1 Store1) shipping(product Product, address string) string {
	for _, p := range s1.products {
		if product == p {
			return "La direccion es " + address
		}
	}
	return "No existe el producto que se quiere entregar"
}

func (s2 Store2) shipping(product Product, address string) string {
	for _, p := range s2.products {
		if product == p {
			return "La direccion es " + address
		}
	}
	return "No existe el producto que se quiere entregar"
}

func (s Small) price() float64 {
	return s.basePrice
}

func (m Medium) price() float64 {
	return m.basePrice + (m.basePrice * 0.03)
}

func (l Large) price() float64 {
	return l.basePrice + (l.basePrice * 0.06) + 2500
}

func calculateFinalPrice(productType string, price float64, productName string) Product {
	switch productType {
	case SMALL:
		return Small{name: productName, basePrice: price}
	case MEDIUM:
		return Medium{name: productName, basePrice: price}
	case LARGE:
		return Large{name: productName, basePrice: price}
	}
	return nil
}

func main() {

	store1 := Store1{name: "Exito", color: "Amarillo", products: []Product{}}
	store2 := Store2{name: "Carulla", products: []Product{}}

	smallProduct1 := store1.addProduct(SMALL, 4200, "Cepillo")
	addressSmallProduct1 := store1.shipping(smallProduct1, "Carrera 23 #50-15")
	fmt.Println(addressSmallProduct1)
	smallProduct2 := store1.addProduct(SMALL, 2800, "Pasta de dientes")
	addressSmallProduct2 := store1.shipping(smallProduct2, "Carrera 25 #81-42")
	fmt.Println(addressSmallProduct2)
	largeProduct1 := store1.addProduct(LARGE, 3200, "Mesa")
	addressLargeProduct1 := store1.shipping(largeProduct1, "Carrera 87 #13-45")
	fmt.Println(addressLargeProduct1)
	fmt.Printf("El valor total de la suma de precios de productos es $%.2f\n", store1.storePrice())
	fmt.Println(store1.products)
	fmt.Println()

	mediumProduct1 := store2.addProduct(MEDIUM, 3200, "Computador")
	addressMediumProduct1 := store2.shipping(mediumProduct1, "Carrera 123 #35-15")
	fmt.Println(addressMediumProduct1)
	smallProduct3 := store2.addProduct(SMALL, 5600, "Celular")
	addressSmallProduct3 := store2.shipping(smallProduct3, "Carrera 98 12-56")
	fmt.Println(addressSmallProduct3)
	mediumProduct2 := store2.addProduct(MEDIUM, 6100, "Sarten")
	addressMediumProduct2 := store2.shipping(mediumProduct2, "Carrera 113 #23-12")
	fmt.Println(addressMediumProduct2)
	largeProduct2 := store2.addProduct(LARGE, 8200, "Televisor")
	addressLargeProduct2 := store2.shipping(largeProduct2, "Carrera 89 #34-23")
	fmt.Println(addressLargeProduct2)
	fmt.Printf("El valor total de la suma de precios de productos es $%.2f\n", store2.storePrice())
	fmt.Println(store2.products)

}
