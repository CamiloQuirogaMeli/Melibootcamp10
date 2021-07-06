package main

import (
	"fmt"
)

type Producto struct {
	nombre   string
	precio   float64
	cantidad float64
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados float64
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func SumarProductos(productos []Producto, c chan float64) {
	var precioTotal float64
	for _, pro := range productos {
		precioTotal += pro.precio * float64(pro.cantidad)
	}
	fmt.Println("El precio total es:", precioTotal)

	c <- precioTotal
}

func SumarServicios(servicios []Servicio, c chan float64) {
	var precioTotal float64
	precioTotal = 0.0

	var mediasHoras float64
	for _, ser := range servicios {
		mediasHoras += float64(ser.minutosTrabajados) / 120
		precioTotal += ser.precio * (float64(ser.minutosTrabajados) / 120)
	}
	if mediasHoras <= 1 {
		fmt.Println("Trabajó", mediasHoras*2, "horas, se cobra media hora: ", precioTotal)

	}
	//fmt.Printf("Trabajó un total de", mediasHoras*2, "horas")
	fmt.Println("Se cobra un total de", precioTotal)
	c <- precioTotal
}

func SumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	var precioTotal float64
	precioTotal = 0.0
	for _, man := range mantenimientos {
		precioTotal += man.precio
	}
	fmt.Println("Se cobra un total de", precioTotal)
	c <- precioTotal
}

func main() {
	/*

			Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
		Para ello requieren realizar un programa que se encargue de calcular el precio total de
		Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go
		routines.

		Se requieren 3 estructuras:
		- Productos: nombre, precio, cantidad.
		- Servicios: nombre, precio, minutos trabajados.
		- Mantenimiento: nombre, precio.

		Se requieren 3 funciones:
		- Sumar Productos: recibe un array de producto y devuelve el precio total (precio *
		cantidad).
		- Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media
		hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado
		media hora).
		- Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

		Los 3 se deben ejecutar en paralelo y al final se debe mostrar por pantalla el monto final
		(sumando el total de los 3).

	*/
	c := make(chan float64)
	productos := make([]Producto, 0)
	servicios := make([]Servicio, 0)
	mantenimientos := make([]Mantenimiento, 0)

	var nombre string
	var precio float64
	var cantidad float64
	var minutos float64
	var suma float64

	salir := false
	var opcion int

	for !salir {
		fmt.Println("Digita una opción:")

		fmt.Println("\t 1: Agregar un producto")
		fmt.Println("\t 2: Agregar un servicio")
		fmt.Println("\t 3: Agregar un mantenimiento")
		fmt.Println("\t 4: Mostrar los productos y sumarlos")
		fmt.Println("\t 5: Mostrar los servicios y sumarlos")
		fmt.Println("\t 6: Mostrar los mantenimientos y sumarlos")
		fmt.Println("\t 7: Realizar las sumas simultáneamente")

		fmt.Println("\t 8: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Digita el nombre del producto:")
			fmt.Scanln(&nombre)
			fmt.Println("Digita el precio del producto:")
			fmt.Scanln(&precio)
			fmt.Println("Digita lacantidad del producto:")
			fmt.Scanln(&cantidad)
			productos = append(productos, Producto{
				nombre:   nombre,
				precio:   precio,
				cantidad: cantidad,
			})
		case 2:
			fmt.Println("Digita el nombre del servicio:")
			fmt.Scanln(&nombre)
			fmt.Println("Digita el precio del servicio:")
			fmt.Scanln(&precio)
			fmt.Println("Digita los minutos trabajados del servicio:")
			fmt.Scanln(&minutos)
			servicios = append(servicios, Servicio{
				nombre:            nombre,
				precio:            precio,
				minutosTrabajados: minutos,
			})

		case 3:
			fmt.Println("Digita el nombre del mantenimiento:")
			fmt.Scanln(&nombre)
			fmt.Println("Digita el precio del mantenimiento:")
			fmt.Scanln(&precio)

			mantenimientos = append(mantenimientos, Mantenimiento{
				nombre: nombre,
				precio: precio,
			})
		case 4:
			fmt.Println("Los productos guardados hasta el momento son los siguientes:")
			for i, pro := range productos {
				fmt.Println("Producto", i+1)
				fmt.Println("\tNombre", pro.nombre)
				fmt.Println("\tPrecio", pro.precio)
				fmt.Println("\tCantidad", pro.cantidad)
			}

			go SumarProductos(productos, c)
			fmt.Println("Por canal:", <-c)

		case 5:
			fmt.Println("Los servicios guardados hasta el momento son los siguientes:")
			for i, ser := range servicios {
				fmt.Println("Producto", i+1)
				fmt.Println("\tNombre: ", ser.nombre)
				fmt.Println("\tPrecio: ", ser.precio)
				fmt.Println("\tMinutos trabajados", ser.minutosTrabajados)
			}

			go SumarServicios(servicios, c)
			fmt.Println("Por canal: ", <-c)
		case 6:
			fmt.Println("Los mantenimientos guardados hasta el momento son los siguientes:")
			for i, man := range mantenimientos {
				fmt.Println("mantenimientos", i+1)
				fmt.Println("\tNombre", man.nombre)
				fmt.Println("\tPrecio", man.precio)
			}

			go SumarMantenimiento(mantenimientos, c)
			fmt.Println("Por canal: ", <-c)
		case 7:
			fmt.Println("Sumas simultáneas:")
			go SumarProductos(productos, c)
			suma += <-c
			go SumarServicios(servicios, c)
			suma += <-c
			go SumarMantenimiento(mantenimientos, c)
			suma += <-c
			fmt.Println("La suma total es", suma)
		case 8:
			salir = true
		}
	}

}
