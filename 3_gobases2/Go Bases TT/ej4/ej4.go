package main

import (
	"fmt"
	"os"
	"reflect"
)

const (
	CHICO           = "chico"
	MEDIANO         = "mediano"
	GRANDE          = "grande"
	FRAGIL          = "fragil"
	ESPECIAL        = "especial"
	CAPACIDADMAXIMA = 10000000
)

type Flate struct {
	Productos []Producto
	ocupacion float32
}
type Producto interface {
	espacioOcupado() float32
}
type Chico struct {
	tamanio float32
}
type Mediano struct {
	tamanio float32
}
type Grande struct {
	tamanio float32
}
type Fragil struct {
	tamanio float32
}
type Especial struct {
	tamanio float32
}

func (c Chico) espacioOcupado() float32 {
	return c.tamanio
}
func (c Mediano) espacioOcupado() float32 {
	return c.tamanio * 1.05
}
func (c Grande) espacioOcupado() float32 {
	return c.tamanio * 1.2
}
func (f Fragil) espacioOcupado() float32 {
	return f.tamanio * 1.75
}
func (f Especial) espacioOcupado() float32 {
	return f.tamanio
}
func newProducto(tipo string, tam float32) Producto {
	switch tipo {
	case CHICO:
		return Chico{tamanio: tam}
	case MEDIANO:
		return Mediano{tamanio: tam}
	case GRANDE:
		return Grande{tamanio: tam}
	case FRAGIL:
		return Fragil{tamanio: tam}
	case ESPECIAL:
		return Especial{tamanio: tam}
	default:
		return nil
	}
}
func main() {
	flete := Flate{}
	flete.ocupacion = 0.0
	var opcion int
	for opcion != 7 {
		fmt.Println("Agregar producto al flete")
		fmt.Println("1. Producto chico")
		fmt.Println("2. Producto mediano")
		fmt.Println("3. Producto grande")
		fmt.Println("4. Producto fragil")
		fmt.Println("5. Producto especial")
		fmt.Println("6. Ver espacio disponible")
		fmt.Println("7. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			if len(flete.Productos) > 0 && reflect.TypeOf(flete.Productos[len(flete.Productos)-1]) == reflect.TypeOf(Especial{}) {
				fmt.Println("Solo puede agregar productos especiales")
				continue
			}
			var tam float32
			fmt.Println("Tamaño en cm3: ")
			fmt.Scanln(&tam)
			p := newProducto(CHICO, tam)
			if flete.ocupacion < CAPACIDADMAXIMA {
				flete.Productos = append(flete.Productos, p)
				flete.ocupacion += p.espacioOcupado()
			}
		case 2:
			if len(flete.Productos) > 0 && reflect.TypeOf(flete.Productos[len(flete.Productos)-1]) == reflect.TypeOf(Especial{}) {
				fmt.Println("Solo puede agregar productos especiales")
				continue
			}
			var tam float32
			fmt.Println("Tamaño en cm3: ")
			fmt.Scanln(&tam)
			p := newProducto(CHICO, tam)
			if flete.ocupacion < CAPACIDADMAXIMA {
				flete.Productos = append(flete.Productos, p)
				flete.ocupacion += p.espacioOcupado()
			}
		case 3:
			if len(flete.Productos) > 0 && reflect.TypeOf(flete.Productos[len(flete.Productos)-1]) == reflect.TypeOf(Especial{}) {
				fmt.Println("Solo puede agregar productos especiales")
				continue
			}
			var tam float32
			fmt.Println("Tamaño en cm3: ")
			fmt.Scanln(&tam)
			p := newProducto(CHICO, tam)
			if flete.ocupacion < CAPACIDADMAXIMA {
				flete.Productos = append(flete.Productos, p)
				flete.ocupacion += p.espacioOcupado()
			}
		case 4:
			if len(flete.Productos) > 0 && reflect.TypeOf(flete.Productos[len(flete.Productos)-1]) == reflect.TypeOf(Especial{}) {
				fmt.Println("Solo puede agregar productos especiales")
				continue
			}
			var tam float32
			fmt.Println("Tamaño en cm3: ")
			fmt.Scanln(&tam)
			p := newProducto(CHICO, tam)
			if flete.ocupacion < CAPACIDADMAXIMA {
				flete.Productos = append(flete.Productos, p)
				flete.ocupacion += p.espacioOcupado()
			}
		case 5:
			var tam float32
			if reflect.TypeOf(flete.Productos[len(flete.Productos)-1]) == reflect.TypeOf(Especial{}) {
				fmt.Println("Tamaño en cm3: ")
				fmt.Scanln(&tam)
				p := newProducto(ESPECIAL, tam)
				if flete.ocupacion < CAPACIDADMAXIMA {
					flete.Productos = append(flete.Productos, p)
					flete.ocupacion += p.espacioOcupado()
				}
			} else {
				fmt.Println("No puede agregar este tipo de producto")
			}
		case 6:
			fmt.Printf("Espacio disponible: %.2f\n", float32(CAPACIDADMAXIMA-flete.ocupacion))
		case 7:
			os.Exit(1)
		default:
			fmt.Println("--Opción inválida--")
		}

	}
}
