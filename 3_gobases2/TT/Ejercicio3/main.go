package main

import "fmt"

const (
	CHICO   = "CHICO"
	MEDIANO = "MEDIANO"
	GRANDE  = "GRANDE"
)

func nuevaTienda(tipo string, precio float64, direccion string) Ecommerce {
	switch tipo {
	case CHICO:
		return ProdChico{precio, direccion}
	case MEDIANO:
		return ProdMed{precio, direccion}
	case GRANDE:
		return ProdGrande{precio, direccion}
	}
	return nil
}

type Ecommerce interface {
	precio() float64
	envio() string
}

type ProdChico struct {
	pre float64
	dir string
}

func (p ProdChico) precio() float64 {
	return p.pre
}
func (p ProdChico) envio() string {
	return p.dir
}

type ProdMed struct {
	pre float64
	dir string
}

func (p ProdMed) precio() float64 {
	return p.pre * 1.03
}
func (p ProdMed) envio() string {
	return p.dir
}

type ProdGrande struct {
	pre float64
	dir string
}

func (p ProdGrande) precio() float64 {
	return p.pre*1.06 + 2500
}
func (p ProdGrande) envio() string {
	return p.dir
}

func detalles(e Ecommerce) {
	fmt.Println("Precio Producto: ", e.precio())
	fmt.Println("Direccion Cliente: ", e.envio())

}

func main() {
	p1 := nuevaTienda(CHICO, 75.89, "larsen")
	p2 := nuevaTienda(MEDIANO, 1854.9, "correa 5896")
	p3 := nuevaTienda(GRANDE, 85124.5, "ramallo 555")
	p4 := nuevaTienda("hola", 7876.3, "amenabar 8547")

	detalles(p1)
	detalles(p2)
	detalles(p3)

	fmt.Println("-----------")

	fmt.Println("Precio Producto 1: ", p1.precio(), "Direccion :", p1.envio())
	fmt.Println("Precio Producto 2: ", p2.precio(), "Direccion: ", p2.envio())
	fmt.Println("Precio Producto 3: ", p3.precio(), "Direccion: ", p3.envio())
	fmt.Println(p4)
}
