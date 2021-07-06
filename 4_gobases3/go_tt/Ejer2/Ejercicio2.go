/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
- Nuevo producto: recibe nombre y precio, y retorna un producto.
- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el
producto al usuario.
- Borrar productos: recibe un usuario, borra los productos del usuario.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

type Producto struct {
	Nombre   string
	Precio   string
	Cantidad string
}

func nuevoProducto(nombre string, precio string) Producto {

	producto := &Producto{}
	producto.Nombre = nombre
	producto.Precio = precio

	return *producto
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad string) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
}

func borrarProducto(usuario *Usuario) {
	usuario.Productos = usuario.Productos[:0]
}

func main() {
	OPTIONS := `
	1. Nuevo producto
	2. Agregar producto
	3. Eliminar producto
	4. Ver información de usuario y productos
	5. Salir
	`
	scanner := bufio.NewScanner(os.Stdin)
	user := &Usuario{
		Nombre:   "David",
		Apellido: "Pelaez",
		Correo:   "david@xyz.com",
	}
	producto := &Producto{}
	fmt.Println("Bienvenido al programa")

	continous := true
	for continous {
		fmt.Println("Digite el numero de la funcion que desee realizar")
		fmt.Println(OPTIONS)
		var option int
		_, errOption := fmt.Scanf("%d", &option)

		if errOption != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}

		switch option {
		case 1:

			fmt.Println("Nombre del producto")
			scanner.Scan()
			nombre := scanner.Text()
			fmt.Println("Precio del producto")
			scanner.Scan()
			precio := scanner.Text()
			*producto = nuevoProducto(nombre, precio)

		case 2:
			fmt.Println("¿Desea agregar el producto al usuario? (Y o N)")
			scanner.Scan()
			decision := strings.ToUpper(scanner.Text())

			if decision == "Y" {
				fmt.Println("Cual es la cantidad del producto")
				scanner.Scan()
				agregarProducto(user, producto, scanner.Text())
			} else if decision == "N" {
				fmt.Println("No se agrego el producto")
			} else {
				fmt.Println("Bad input")
				os.Exit(0)
			}
		case 3:
			borrarProducto(user)
			fmt.Println("Se han eliminado todos los productos")
		case 4:
			fmt.Println(user.Nombre, user.Apellido, user.Correo)
			fmt.Println("Los productos del usuario son:")
			for i, product := range user.Productos {
				fmt.Println(i, "Nombre:", product.Nombre, "Precio:", product.Precio, "Cantidad:", product.Cantidad)
			}
		case 5:
			fmt.Println("Gracias por utilizar el programa")
			os.Exit(0)
		default:
			fmt.Println("La opcion no existe")
		}
	}

}
