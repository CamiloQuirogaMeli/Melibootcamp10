package main

import "fmt"

func main() {
	fmt.Print("Ingrese ejercicio a resolver (1 o 2 o 3 o 4):\n")
	var ejercicioAresolver int
	fmt.Scanln(&ejercicioAresolver)
	switch ejercicioAresolver {
	case 1:
		fmt.Print("Ingrese nombre de Alumno:\n")
		var fnameAlum string
		fmt.Scanln(&fnameAlum)
		fmt.Print("Ingrese apellido de Alumno:\n")
		var lnameAlum string
		fmt.Scanln(&lnameAlum)
		fmt.Print("Ingrese DNI de Alumno:\n")
		var dniAlum int
		fmt.Scanln(&dniAlum)
		fmt.Print("Ingrese fecha de ingreso de Alumno:\n")
		var fechaIngresAlum string
		fmt.Scanln(&fechaIngresAlum)

		p1 := Estudiante{fnameAlum, lnameAlum, dniAlum, fechaIngresAlum}

		p1.detalle()

	case 2:

		a := []float64{3, 4}
		b := []float64{3, 4}
		matrix := [][]float64{a, b}
		var m Matrix
		m.set(matrix)
		m.print()
	case 3:
		fmt.Print("Ingrese nombre de Tienda:\n")
		var tienda string
		fmt.Scanln(&tienda)
		fmt.Print("Ingrese tipo de Producto:\n")
		var tipoProducto string
		fmt.Scanln(&tipoProducto)
		fmt.Print("Ingrese valor de Producto:\n")
		var valorProducto float64
		fmt.Scanln(&valorProducto)
		fmt.Print("Ingrese ubicación de envío:\n")
		var ubic string
		fmt.Scanln(&ubic)

		/*
			p := producto{tipo: tipoProducto, valorProd: valorProducto}
			detalle(p)
		*/
		//p1 := nuevaTienda(tipoProducto, valorProducto)
		if tienda == "MELI" {
			fmt.Println("La tienda a la que está accediendo es: ", tienda)
			t1 := MELI{NombreTienda: tienda, prod: Producto{tipoProducto, valorProducto, ubic}}
			nuevaTienda(t1.prod.tipo, t1.prod.valorProd, t1.prod.ubicacionEnvio)
			detalle(t1.prod)
		} else if tienda == "AMAZON" {
			fmt.Println("La tienda a la que está accediendo es: ", tienda)
			t2 := AMAZON{NombreTienda: tienda, prod: Producto{tipoProducto, valorProducto, ubic}}
			nuevaTienda(t2.prod.tipo, t2.prod.valorProd, t2.prod.ubicacionEnvio)
			detalle(t2.prod)
		}

	case 4:

		p1 := fabrica(prod_chico, 100000)
		p2 := fabrica(prod_mediano, 250000)
		p3 := fabrica(prod_grande, 3500000)
		p4 := fabrica(prod_fragil, 4000000)
		p5 := fabrica(prod_especial, 52000)
		p6 := fabrica(prod_especial, 60000)
		p7 := fabrica(prod_especial, 123456789)
		p8 := fabrica(prod_especial, 123456789)
		p9 := fabrica(prod_especial, 123456789)

		f1 := Flete{}
		f1.agregarProducto(p1, p2, p3, p4, p5, p6, p7, p8, p9)

		cantOtros, cantEsp := f1.calcularEnvio()

		fmt.Println("La cantidad de viajes especiales será de: ", cantEsp)
		fmt.Println("La cantidad de viajes de productos no especiales será de: ", cantOtros)
		fmt.Println("La cantidad total de viajes es de: ", cantEsp+cantOtros)
	case 5:
		break

	}
}
