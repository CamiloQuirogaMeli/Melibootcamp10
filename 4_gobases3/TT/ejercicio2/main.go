package main

import (
	"errors"
	"fmt"
)

type User struct {
	name     string
	age      int
	email    string
	password string
	products []Product
}

type Product struct {
	name   string
	price  float64
	amount int
}

func newProduct() Product {
	var name string
	var price float64
	var amount int

	fmt.Println("Ingrese nombre")
	fmt.Scan(&name)

	fmt.Println("Ingrese precio")
	fmt.Scan(&price)

	fmt.Println("Ingrese cantidad")
	fmt.Scan(&amount)

	product := Product{name, price, amount}
	return product
}

func newUser() User {
	var name, email, password string
	var age int

	fmt.Println("Ingrese nombre")
	fmt.Scan(&name)

	fmt.Println("Ingrese el correo")
	fmt.Scan(&email)

	fmt.Println("Ingrese la edad")
	fmt.Scan(&age)

	fmt.Println("Ingrese el password")
	fmt.Scan(&password)

	user := User{name, age, email, password, []Product{}}
	return user
}

func findUser(name string, users map[string]User) (User, error) {

	user, exists := users[name]
	if exists {
		return user, nil
	} else {
		return user, errors.New("¡¡ERROR!! el usuario no existe")
	}
}
func findProduct(name string, products map[string]Product) (Product, error) {

	product, exists := products[name]
	if exists {
		return product, nil
	} else {
		return product, errors.New("¡¡ERROR!! el producto no existe")
	}
}

func addProduct(user *User, product *Product, amount int) {

	for i := 0; i < amount; i++ {
		user.products = append(user.products, *product)
	}
}

func deleteProduct(user *User) {
	user.products = []Product{}
}

func main() {
	var option int
	var aux int = 1
	users := make(map[string]User)
	products := make(map[string]Product)
	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Crear producto\n2. Crear usuario\n3. Agregar producto a un usuario\n4. Borrar los productos de un usuario \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			product := newProduct()
			products[product.name] = product

			fmt.Println("El producto ha sido creado exitosamente")
		case 2:
			user := newUser()
			users[user.name] = user
			fmt.Println("El usuario ha sido creado exitosamente")
		case 3:
			var productName, userName string
			var amount int
			fmt.Println("Ingrese el nombre del producto que desea agregar")
			fmt.Scan(&productName)

			product, err := findProduct(productName, products)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Ingrese el nombre del usuario al cual le va a agregar el producto")
			fmt.Scan(&userName)
			user, err := findUser(userName, users)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Ingrese el numero de productos que desea agregar")
			fmt.Scan(&amount)

			addProduct(&user, &product, amount)

			fmt.Println("El producto ha sido agregado con exito")
		case 4:
			var userName string
			fmt.Println("Ingrese el nombre del usuario al cual le va a eliminar los productos")
			fmt.Scan(&userName)
			user, err := findUser(userName, users)
			if err != nil {
				fmt.Println(err)
				break
			}
			deleteProduct(&user)

			fmt.Println("El usuario ya no tiene productos")

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}
}
