package main

import (
	"fmt"
)

func main() {

	for {
		fmt.Println("Ingrese el precio de la prenda")
		fmt.Println("---------------------")
		fmt.Print("-> ")
		var i int
		fmt.Scan(&i)

		fmt.Println("Ingrese el descuento de la prenda")
		fmt.Println("---------------------")
		fmt.Print("-> ")
		var j int
		fmt.Scan(&j)

		var res = (j * i) / 100
		res = i - res
		fmt.Println("La prenda que le interesa tiene un costo de $", res)

	}
}
