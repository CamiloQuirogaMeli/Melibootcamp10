package main

import "fmt"

type Producto struct{
	costo float64
	direccion string
}

type Chico struct {
	p Producto
}

type Mediano struct{
	p Producto
	adicional float64
}

type Grande struct{
	p Producto
	adicional float64
	envio uint
}

func (p Chico) Precio() float64{
	return p.p.costo
}

func (p Chico) Envio() string{
	return p.p.direccion
}

func (p Mediano) Precio() float64{
	p.p.costo += p.p.costo * p.adicional
	return p.p.costo
}

func (p Mediano) Envio() string{
	return p.p.direccion
}

func (p Grande) Precio() float64{
	p.p.costo += (p.p.costo * p.adicional) + float64(p.envio)
	return p.p.costo
}

func (p Grande) Envio() string{
	return p.p.direccion
}

type Tienda struct{
	tamanioProd string
	direccionEnvio string
}

type Ecommerce interface{
Precio() float64
Envio() string
}

func nuevaTienda (t Tienda) Ecommerce{
	switch t.tamanioProd{
	case "Chico":
		return Chico{Producto{costo: 1200, direccion: t.direccionEnvio}}
	case "Mediano":
		return Mediano{Producto{costo: 3000, direccion: t.direccionEnvio}, 0.03}
	case "Grande":
		return Grande{p:Producto{costo: 5000, direccion: t.direccionEnvio}, adicional: 0.06, envio: 2500}
	}
	return nil
}

func main(){
	t1 := Tienda{tamanioProd: "Chico", direccionEnvio: "Tucuman 355"}
	pt1 := nuevaTienda(t1)
	t2 := Tienda{tamanioProd: "Mediano", direccionEnvio: "Roma 2234"}
	pt2 := nuevaTienda(t2)
	t3 := Tienda{tamanioProd: "Grande", direccionEnvio: "Paul Groussac 9972"}
	pt3 := nuevaTienda(t3)

	fmt.Println("El precio del producto de la tienda 1 es de: $", pt1.Precio())
	fmt.Println("Direccion a enviar: ", pt1.Envio())
	fmt.Println("\nEl precio del producto de la tienda 2 es de: $", pt2.Precio())
	fmt.Println("Direccion a enviar: ", pt2.Envio())
	fmt.Println("\nEl precio del producto de la tienda 3 es de: $", pt3.Precio())
	fmt.Println("Direccion a enviar: ", pt3.Envio())
}

/*Un Ecommerce necesita realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Chico, Mediano y Grande. (Se espera que sean
muchos más)

Y los costos adicionales son:
- Chico: solo tiene el costo del producto
- Mediano: el precio del producto + un %3 de mantenerlo en la tienda
- Grande: el precio del producto + un %6 de mantenerlo en la tienda y adicional a eso
$2500 de flete.

Requerimientos:
Crear dos estructuras “tiendaUno” y “tiendaDos” (Atributos de la estructura y nombre de la misma a elección).
Crear una interface “Ecommerce” que tenga los métodos “Precio” y “Envio”.
Se requiere una función “nuevaTienda” que reciba el tipo de producto. Luego retorne una interface “Ecommerce”
Interface Ecommerce:
 - El método “Precio” debe retornar el precio total en base al costo del producto y los adicionales si los hubiera.
 - El método “Envio” debe retornar la dirección de entrega especificada por el cliente.
*/