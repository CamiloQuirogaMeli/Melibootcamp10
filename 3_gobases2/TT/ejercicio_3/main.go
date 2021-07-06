package main

import (
	"fmt"
	ecommerce "github.com/tomastropea/TT/ejercicio_3/ecommerce"
)

func main() {

	fmt.Println("* Caso producto chico")
	store1, _ := ecommerce.NewStore(ecommerce.SMALL_PRODUCT, ecommerce.SPONGE, "Calle Falsa 123")
	fmt.Printf("\tPrecio final de tienda: $%f\n", store1.Price())
	fmt.Printf("\tLa tienda envia el producto a la direccion: %s\n", store1.ShippingAddress())

	fmt.Println("* Caso producto mediano")
	store2, _ := ecommerce.NewStore(ecommerce.MEDIUM_PRODUCT, ecommerce.SPONGE, "Calle True 321")
	fmt.Printf("\tPrecio final de tienda: $%f\n", store2.Price())
	fmt.Printf("\tLa tienda envia el producto a la direccion: %s\n", store2.ShippingAddress())

	fmt.Println("* Caso producto grande")
	store3, _ := ecommerce.NewStore(ecommerce.LARGE_PRODUCT, ecommerce.SPONGE, "Calle Cuantica 0.5")
	fmt.Printf("\tPrecio final de tienda: $%f\n", store3.Price())
	fmt.Printf("\tLa tienda envia el producto a la direccion: %s\n", store3.ShippingAddress())

	fmt.Println("* Caso store invalido")
	_, err := ecommerce.NewStore(ecommerce.SMALL_PRODUCT, "invalid store", "")
	fmt.Println("\tError:", err)
}
