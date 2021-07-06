package main

import "fmt"

//Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
//La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
//Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un adicional al momento de ser enviado:
//Chico: Ningún adicional.
//Mediano: Requiere un %5 más de espacio
//Grande: Requiere un %20 más de espacio
//Frágil: Requiere un %75 más de espacio
//Especial: Sólo puede ser enviado con productos especiales

//Para ello requerimos que los productos guarden el tamaño y tengan un metodo TamanoTotal que nos devuelva el el espacio en cm3 que requerimos para ser enviado.

//Y una estructura Flete que tenga los métodos:
//- AgregarProducto: agregar producto al flete
//- CalcularEnvios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede cargar un total de 10.000.000 cm3 por envío

type producto interface {
	TamanoTotal() float64
}

type Producto struct {
	Tamano float64
	Adicionales float64
	Especial bool
}

type ProductoChico struct {
	Producto
}

func (p ProductoChico) TamanoTotal() float64{
	return p.Tamano
}

type ProductoMediano struct {
	Producto
}

func (p ProductoMediano) TamanoTotal() float64{
	return p.Tamano + p.Adicionales
}

type ProductoGrande struct {
	Producto
}

func (p ProductoGrande) TamanoTotal() float64{
	return p.Tamano + p.Adicionales
}

type ProductoEspecial struct {
	Producto
}

func (p ProductoEspecial) TamanoTotal() float64{
	return p.Tamano
}

type ProductoFragil struct {
	Producto
}

func (p ProductoFragil) TamanoTotal() float64{
	return p.Tamano + p.Adicionales
}

type Flete struct {
	ProductosFlete []Producto
	CantEnvios int
}

func (f Flete) AgregarProducto(p Producto) {
	ProductosFlete = append(ProductosFlete, p)
}

func (f Flete) CalcularEnvios() int{
	cant := 0
	cantEs := 0
	sum := 0
	sumEs := 0

	for i := 0; i < len(f.ProductosFlete); i++ {

		if f.ProductosFlete[i].Especial {
			if (sumEs + f.ProductosFlete[i].TamanoTotal()) > 10000000 {
				cantEs++
				sumEs = 0
			} else {
				sumEs += f.ProductosFlete[i].TamanoTotal()
			}
		} else { 
			if (sum + f.ProductosFlete[i].TamanoTotal()) > 10000000 {
				cant++
				sum = 0
			} else {
				sum += f.ProductosFlete[i].TamanoTotal()
			}
		}

		if i == (len(f.ProductosFlete) - 1) {
			if sum > 0 {
				cant++
			}

			if sumEs > 0 {
				cantEs++
			}
		}
	}

	f.CantEnvios = cant + cantEs

	return cant + cantEs
}

const (
	chico = "C"
	mediano = "M"
	grande = "G"
	especial = "E"
	fragil = "F"
)

func main() {
	c := factory(chico, 50000)
	m := factory(mediano, 87000)
	g := factory(grande, 120000)
	e := factory(especial, 5000)
	f := factory(fragil, 200)

	flete := Flete {}

	flete.AgregarProducto(c)
	flete.AgregarProducto(m)
	flete.AgregarProducto(g)
	flete.AgregarProducto(e)
	flete.AgregarProducto(f)

	fmt.Println("La cantidad de envíos de este flete es de:", flete.CalcularEnvios())
}

func factory(tipo string, tam float64) producto {
	switch tipo {
	case chico:
		return ProductoChico{Producto{Tamano: tam, Adicionales: 0, Especial: false},}
	case mediano:
		return ProductoMediano{Producto{Tamano: tam, Adicionales: tam * 0.05, Especial: false},}
	case grande:
		return ProductoGrande{Producto{Tamano: tam, Adicionales: tam * 0.2, Especial: false},}
	case fragil:
		return ProductoFragil{Producto{Tamano: tam, Adicionales: tam * 0.75, Especial: false},}
	case especial:
		return ProductoEspecial{Producto{Tamano: tam, Adicionales: 0, Especial: true},}
	}

	return nil
}