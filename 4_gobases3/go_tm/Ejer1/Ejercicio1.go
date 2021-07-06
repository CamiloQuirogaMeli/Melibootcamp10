/*
Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información
de productos comprados, separados por punto y coma.
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Product struct {
	ID       string
	Price    string
	Quantity string
}

func main() {
	fmt.Println("Bienvenidos a su progrma para hacer la lista de compras")
	fmt.Println("¿Cuantos productos comprara?")
	var quantity int
	_, errQuantity := fmt.Scanf("%d", &quantity)

	if errQuantity != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	LimpiarTexto()
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= quantity; i++ {
		product := Product{}
		fmt.Println("Producto #", i)
		fmt.Println("ID del producto: ")
		scanner.Scan()
		product.ID = scanner.Text()
		fmt.Println("precio del producto: ")
		scanner.Scan()
		product.Price = scanner.Text()
		fmt.Println("cantidad del producto: ")
		scanner.Scan()
		product.Quantity = scanner.Text()

		GuardarProducto(product)

	}

	fmt.Println("Su lista ha sido almacenada exitosamente, por favor revisar el documento 'products.txt'")

}

func LimpiarTexto() {
	f, err := os.OpenFile("../products.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	f.Truncate(0)
	defer f.Close()
}

func GuardarProducto(producto Product) {
	f, err := os.OpenFile("../products.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ocurrio un error con los permisos del archivo")
		os.Exit(0)
	}
	var productString string = "ID: " + producto.ID + " PRECIO: " + producto.Price + " CANTIDAD: " + producto.Quantity + "; "

	if _, err := f.WriteString(productString); err != nil {
		fmt.Println("Ocurrio un error almacenando el producto")
		os.Exit(0)
	}

}
