package main

const (
	chico   = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type Ecommerce interface {
	Precio() float64
	Envio() string
}

type TiendaUno struct {
	nombre    string
	direccion string
}

type TiendaDos struct {
	nombre    string
	direccion string
}

func Precio() float64 {
	switch expression {
	case condition:

	}
}

func (t1 TiendaUno) Envio() string {
	return t1.direccion
}

func (t2 TiendaDos) Envio() string {
	return t1.direccion
}

func nuevaTienda(producto string) Ecommerce {

}

func main() {

}
