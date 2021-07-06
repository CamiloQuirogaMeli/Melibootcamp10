package main

import "fmt"

/*

Varias tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
- Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de
envío.

Sus costos adicionales son:
- Pequeño: El costo del producto (sin costo adicional)
- Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén
de la tienda.
- Grande: El costo del producto + un 6% por mantenimiento, y un costo adicional por
envío de $2500.
Requerimientos:
- Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre
de la misma a elección).
- Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
- Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne
una interface “Ecommerce”
- Interface Ecommerce:
- El método “Precio” debe retornar el precio total en base al costo del producto y los
adicionales si los hubiera.
- El método “Envio” debe retornar la dirección de entrega especificada por el cliente.

*/

const (
	CHICO   = "chico"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

/*type tiendaUno struct{
	nombre string
}

type tiendaDos struct{
	nombre string
}*/

type tipo1 struct {
	precioProd float64
	direccion  string
}

type tipo2 struct {
	precioProd float64
	direccion  string
}

type tipo3 struct {
	precioProd float64
	direccion  string
}

type tipoNOT struct {
	mensaje string
}

type Ecommerce interface {
	precio() float64
	envio() string
}

func (t tipo1) precio() float64 {
	return t.precioProd
}

func (t tipo2) precio() float64 {
	return t.precioProd * 1.03
}

func (t tipo3) precio() float64 {
	return t.precioProd*1.06 + 2500.0
}

func (t tipoNOT) precio() float64 {
	fmt.Printf(t.mensaje)
	return 1.0
}

func (t tipo1) envio() string {
	return t.direccion
}

func (t tipo2) envio() string {
	return t.direccion
}

func (t tipo3) envio() string {
	return t.direccion
}

func (t tipoNOT) envio() string {
	return t.mensaje
}

func nuevaTienda(tipo string, precio float64, dir string) Ecommerce {

	switch tipo {
	case CHICO:
		return tipo1{precioProd: precio, direccion: dir}
	case MEDIANO:
		return tipo2{precioProd: precio, direccion: dir}
	case GRANDE:
		return tipo3{precioProd: precio, direccion: dir}
	default:
		return tipoNOT{mensaje: "tipo de paquete no valido"}

	}
}

func main() {

	prodChico := nuevaTienda("chico", 5400.0, "Francia 2357")
	fmt.Println("Precio del producto: ", prodChico.precio())
	fmt.Println("Direccion de envio: ", prodChico.envio())

	prodMediano := nuevaTienda("mediano", 10000.0, "Urquiza 2714")
	fmt.Println("Precio del producto: ", prodMediano.precio())
	fmt.Println("Direccion de envio: ", prodMediano.envio())

	prodGrande := nuevaTienda("grande", 15000.0, "San martín 345")
	fmt.Println("Precio del producto: ", prodGrande.precio())
	fmt.Println("Direccion de envio: ", prodGrande.envio())

}
