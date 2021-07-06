package main

type User struct {
	Cart	[]Product	`json: carrito`
}

type Product struct {
	Name	string	`json: "nombre"`
	Price 	float64 `json: "precio"`
}

func NewProduct(name string, price float64) Product {
	return Product{name, price}
}

func AddProductsToUser(user *User, product *Product, quantity int) {
	for i := 0; i < quantity; i++ {
		user.Cart = append(user.Cart, *product)
	}
}

func RemoveProductsFromUser(user *User) {
	user.Cart = user.Cart[:0]
}
