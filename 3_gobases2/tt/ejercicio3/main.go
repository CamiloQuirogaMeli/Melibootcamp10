package main

// Dejo este ejercicio incompleto ya que me resulta confuso comprender el enunciado
// y abstraer el problema
type Product struct {
	size         string
	cost         float64
	costBySize   float64 // in percentage
	extraCost    float64
	shippingAddr string
}

type Shop struct {
	products []Product
}

func (p *Product) Price() {

}

func (p *Product) Shipping() {

}

type Ecommerce interface {
	Price() float64
	Shipping() string
}

func main() {

}
