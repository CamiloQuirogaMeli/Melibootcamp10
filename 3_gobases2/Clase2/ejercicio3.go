package main

import (
	"fmt"
)

type Ecommerce interface {
	Precio() float64
	Envio() (string, float64)
}

type tiendaUno struct {
	//Ideal para tiendas pequeñas
	tipoProducto   string
	direccionEnvio string
	costoProducto  float64
}

func (t *tiendaUno) Set(tipoProducto string, direccionEnvio string, costoProducto float64) {
	t.tipoProducto = tipoProducto
	t.direccionEnvio = direccionEnvio
	t.costoProducto = costoProducto
}

func (tu tiendaUno) Precio() float64 {
	precio := tu.costoProducto
	fmt.Println("Como tu tienda es pequeña, no tendrás costos adicionales en tu producto")
	return precio
}

func (tu tiendaUno) Envio() (string, float64) {
	envio := tu.direccionEnvio
	fmt.Println("Como tu tienda es pequeña, no tendrás costos de envio adicionales")
	return envio, 0
}

type tiendaDos struct {
	//Ideal para tiendas medianas
	tipoProducto             string
	direccionEnvio           string
	costoProducto            float64
	porcentajeCostoAdicional float64
}

func (t *tiendaDos) Set(tipoProducto string, direccionEnvio string, costoProducto float64, porcentajeCostoAdicional float64) {
	t.tipoProducto = tipoProducto
	t.direccionEnvio = direccionEnvio
	t.costoProducto = costoProducto
	t.porcentajeCostoAdicional = porcentajeCostoAdicional
}

func (td tiendaDos) Precio() float64 {
	precio := td.costoProducto
	fmt.Println("Como tu tienda es mediana, tendrás un ", td.porcentajeCostoAdicional, "% por mantenerlo en existencia en el almacén")
	return (precio * td.porcentajeCostoAdicional) / 100
}

func (td tiendaDos) Envio() (string, float64) {
	envio := td.direccionEnvio
	fmt.Println("Como tu tienda es mediana, no tendrás costos de envio adicionales")
	return envio, 0
}

type tiendaTres struct {
	//Ideal para tiendas grandes
	tipoProducto             string
	direccionEnvio           string
	costoProducto            float64
	porcentajeCostoAdicional float64
	costoAdicionalEnvio      float64
}

func (t *tiendaTres) Set(tipoProducto string, direccionEnvio string, costoProducto float64, porcentajeCostoAdicional float64, costoAdicionalEnvio float64) {
	t.tipoProducto = tipoProducto
	t.direccionEnvio = direccionEnvio
	t.costoProducto = costoProducto
	t.porcentajeCostoAdicional = porcentajeCostoAdicional
	t.costoAdicionalEnvio = costoAdicionalEnvio
}

func (tt tiendaTres) Precio() float64 {
	precio := tt.costoProducto
	fmt.Println("Como tu tienda es grande, tendrás un ", tt.porcentajeCostoAdicional, "% por mantenimiento")
	return (precio * tt.porcentajeCostoAdicional) / 100
}

func (tt tiendaTres) Envio() (string, float64) {
	envio := tt.direccionEnvio
	fmt.Println("Como tu tienda es grande, tendrás un costo adicional porenvío de $2500")
	return envio, 2500.0
}

func nuevaTienda(e Ecommerce) {

	fmt.Println("Tienda creada")
	fmt.Println("Los costos de tu tienda son los siguientes", e.Precio())
	direccion, envio := e.Envio()
	fmt.Println("Los costos de envio son:", envio, "y la direccion de envio es:", direccion)
}

func main() {
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
	var tipoProducto string
	var precioProducto float64
	numeroDeTiendas := 3
	var direccion string
	tiendasPequeñas := make([]tiendaUno, 0)
	tp := 0
	tiendasMedianas := make([]tiendaDos, 0)
	tm := 0
	tiendasGrandes := make([]tiendaTres, 0)
	tg := 0

	for tienda := 0; tienda < numeroDeTiendas; tienda++ {
		fmt.Println("Digita tu dirección")
		fmt.Scanln(&direccion)

		fmt.Println("Digita el tipo de producto de la tienda", tienda+1, ":")
		fmt.Println("\t p: Pequeño")
		fmt.Println("\t m: Mediano")
		fmt.Println("\t g: Grande")
		fmt.Scanln(&tipoProducto)
		for tipoProducto != "p" && tipoProducto != "m" && tipoProducto != "g" {
			fmt.Println("Opción incorrecta")
			fmt.Println("Digita el tipo de producto de la tienda", tienda+1, ":")
			fmt.Println("\t p: Pequeño")
			fmt.Println("\t m: Mediano")
			fmt.Println("\t g: Grande")
			fmt.Scanln(&tipoProducto)
		}

		fmt.Println("Digita el precio del producto de la tienda", tienda+1, ":")
		fmt.Scanln(&precioProducto)

		for precioProducto <= 0.0 {
			fmt.Println("Opción incorrecta")
			fmt.Println("Digita el precio del producto de la tienda", tienda+1, ":")

			fmt.Scanln(&precioProducto)
		}
		switch tipoProducto {
		case "p":
			tiendasPequeñas = append(tiendasPequeñas, tiendaUno{})
			tiendasPequeñas[tp].Set(tipoProducto, direccion, precioProducto)
			tp++
		case "m":
			tiendasMedianas = append(tiendasMedianas, tiendaDos{})
			tiendasMedianas[tm].Set(tipoProducto, direccion, precioProducto, 3.0)
			tm++
		case "g":
			tiendasGrandes = append(tiendasGrandes, tiendaTres{})
			tiendasGrandes[tg].Set(tipoProducto, direccion, precioProducto, 6.0, 2500.0)
			tg++
		}

		fmt.Println("\n--------------")
		fmt.Println("Tienda creada")
		fmt.Println("--------------\n")

	}

	fmt.Println("Las tiendas creadas y sus detalles son los siguientes:")
	fmt.Println("Tiendas pequeñas:")
	for i, t := range tiendasPequeñas {
		fmt.Println("Tienda", i+1, ":")
		nuevaTienda(t)
	}
	fmt.Println("Tiendas medianas:")
	for i, t := range tiendasMedianas {
		fmt.Println("Tienda", i+1, ":")
		nuevaTienda(t)
	}
	fmt.Println("Tiendas grandes:")
	for i, t := range tiendasGrandes {
		fmt.Println("Tienda", i+1, ":")
		nuevaTienda(t)
	}

}
