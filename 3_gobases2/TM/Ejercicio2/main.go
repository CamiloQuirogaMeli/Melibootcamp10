package main

import (
	"fmt"
	"os"
)

func main() {
	Ejercicio2()
	Ejercicio2ConEllipsis(10, 8, 6, 7, 4, 9)
}

func Ejercicio2() {
	fmt.Println("Ejercicio2")
	var calificaciones = []uint8{10, 8, 6, 7, 6, 9}
	var sumaCalificaciones float64 = 0.00

	for _, c := range calificaciones {
		if c <= 10 {
			sumaCalificaciones += float64(c)
		} else {
			fmt.Println("Al menos un valor de la calificacion es incorecto")
			os.Exit(0)
		}
	}
	fmt.Printf("El promedio de las notas del alumno es de: ")
	fmt.Printf("%.2f \n", sumaCalificaciones/float64(len(calificaciones)))
	fmt.Println("========================")
}

func Ejercicio2ConEllipsis(calificacionesEllipsis ...uint8) {
	fmt.Println("Ejercicio2ConEllipsis")
	var sumaCalificacionesEllipsis float64 = 0.00

	for _, cEllipsis := range calificacionesEllipsis {
		if cEllipsis <= 10 {
			sumaCalificacionesEllipsis += float64(cEllipsis)
		} else {
			fmt.Println("Al menos un valor de la calificacion es incorecto")
			os.Exit(0)
		}
	}
	fmt.Printf("El promedio de las notas del alumno es de: ")
	fmt.Printf("%.2f \n", sumaCalificacionesEllipsis/float64(len(calificacionesEllipsis)))
}
