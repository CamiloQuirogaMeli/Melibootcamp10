package main

import "fmt"

const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type Chico struct {
	valor            float64
	direccionCliente string
}
type Mediano struct {
	valor            float64
	direccionCliente string
}
type Grande struct {
	valor            float64
	direccionCliente string
}

type Ecommerce interface {
	precio() float64
	envio() string
}

func (c Chico) envio() string {
	return c.direccionCliente
}

func (c Chico) precio() float64 {
	return c.valor
}

func (m Mediano) precio() float64 {
	return m.valor * 1.03
}

func (m Mediano) envio() string {
	return m.direccionCliente
}

func (g Grande) precio() float64 {
	return (g.valor * 1.06) + 2500
}

func (g Grande) envio() string {
	return g.direccionCliente
}

func nuevaTienda(prodType string, p float64, d string) Ecommerce {
	switch prodType {
	case CHICO:
		return Chico{valor: p, direccionCliente: d}
	case MEDIANO:
		return Mediano{valor: p, direccionCliente: d}
	case GRANDE:
		return Grande{valor: p, direccionCliente: d}
	default:
		return nil
	}
}

func main() {
	tiendaUno := nuevaTienda(MEDIANO, 1000.0, "Gaucho Cruz 1205")
	tiendaDos := nuevaTienda(GRANDE, 1510.50, "Congreso 483")

	fmt.Println("La tienda uno enviara a ", tiendaUno.envio(), " un pedido por ", tiendaUno.precio(), " pesos")
	fmt.Println("La tienda dos enviara a ", tiendaDos.envio(), " un pedido por ", tiendaDos.precio(), " pesos")
}
