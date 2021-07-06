package main

import (
	"fmt"
	product "github.com/tomastropea/TT/ejercicio_4/product"
)

func main() {

	smallProduct, _ := product.ProductFactory(product.SMALL_PRODUCT, 10000000)
	mediumProduct, _ := product.ProductFactory(product.MEDIUM_PRODUCT, 2000000)
	largeProduct, _ := product.ProductFactory(product.LARGE_PRODUCT, 1000000)
	fragileProduct, _ := product.ProductFactory(product.FRAGILE_PRODUCT, 1000000)
	specialProduct1, _ := product.ProductFactory(product.SPECIAL_PRODUCT, 5000000)
	specialProduct2, _ := product.ProductFactory(product.SPECIAL_PRODUCT, 8000000)
	
	fmt.Println("Espacio ocupado por producto chico:", smallProduct.TotalSize())
	fmt.Println("Espacio ocupado por producto mediano:", mediumProduct.TotalSize())
	fmt.Println("Espacio ocupado por producto largo:", largeProduct.TotalSize())
	fmt.Println("Espacio ocupado por producto fragil:", fragileProduct.TotalSize())
	fmt.Println("Espacio ocupado por producto especial 2:", specialProduct1.TotalSize())
	fmt.Println("Espacio ocupado por producto especial 1:", specialProduct2.TotalSize())

	products := []product.Product{
		smallProduct,
		mediumProduct,
		largeProduct,
		fragileProduct,
		specialProduct1,
		specialProduct2,
	}
	charter := Charter{}
	for _, item := range products {
		charter.AddProduct(item)
	}
	fmt.Println("Cantidad de envios necesarios:", charter.RequiredShippings())
}