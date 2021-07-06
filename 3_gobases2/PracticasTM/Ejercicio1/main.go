package main

import (
	"fmt"
)

func main() {
	var salario float64 = 160000
	imp := impuesto(salario)
	fmt.Println("El impuesto que debe pagar es de: ", imp)
}

func impuesto(i float64) float64 {
	if i > 50000 && i <= 150000 {
		return i * 17 / 100
	} else if i > 150000 {
		return i * 27 / 100
	} else {
		return i * 0
	}

}
