package main

import "fmt"

type product interface {
	getTotalSize() float64
}

type smallProduct struct {
	size float64
}
type mediumProduct struct {
	size float64
}
type largeProduct struct {
	size float64
}
type fragileProduct struct {
	size float64
}
type specialProduct struct {
	size float64
}

func (p smallProduct) getTotalSize() float64 {
	return p.size
}
func (p mediumProduct) getTotalSize() float64 {
	return p.size * 1.05
}
func (p largeProduct) getTotalSize() float64 {
	return p.size * 1.2
}

func (p fragileProduct) getTotalSize() float64 {
	return p.size * 1.75
}
func (p specialProduct) getTotalSize() float64 {
	return p.size
}

type freight struct {
	cargo []product
}

func (f *freight) addProduct(product product) {

	f.cargo = append(f.cargo, product)

}

func (f freight) calcShipment() int {
	var normalCargoTotalSize, specialCargoTotalSize float64

	for _, prod := range f.cargo {

		switch prod.(type) {
		case specialProduct:
			specialCargoTotalSize += prod.getTotalSize()
		default:
			normalCargoTotalSize += prod.getTotalSize()
		}

	}

	normalShipments := int(normalCargoTotalSize/float64(10000000)) + 1
	specialShipments := int(specialCargoTotalSize/float64(10000000)) + 1

	return normalShipments + specialShipments
}

func main() {

	freight := freight{}

	p1 := smallProduct{1000.00}
	p2 := mediumProduct{1000.00}
	p3 := largeProduct{1000.00}
	p4 := fragileProduct{1000.00}
	ps := specialProduct{1000.00}

	freight.addProduct(p1)
	freight.addProduct(p2)
	freight.addProduct(p3)
	freight.addProduct(p4)
	freight.addProduct(ps)

	fmt.Println("Shipments:", freight.calcShipment())

}
