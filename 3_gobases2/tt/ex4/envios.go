package main

import "fmt"

const (
	SMALL   = "SMALL"
	MEDIUM  = "MEDIUM"
	BIG     = "BIG"
	FRAGILE = "FRAGILE"
	SPECIAL = "SPECIAL"
)

type Product struct {
	size float64
}

type Freigth struct {
	products  int
	shipments int
	capacity  float64
}

func (p Product) totalSize() float64 {
	return 0
}

func (f Freigth) addProduct() int {
	return 0
}

func (f Freigth) calcShip() float64 {
	return 0
}

func main() {
	fmt.Println("vim-go")
	p := Produ
}
