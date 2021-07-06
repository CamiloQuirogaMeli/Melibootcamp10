package ejercicios

import "fmt"

func TercerEjercicio() {
	var (
		age         = 23
		employee    = true
		salary      = 95000
		workingTime = 72
	)

	if age > 22 && employee && workingTime > 12 {
		if salary > 100000 {
			fmt.Println("Su prestamo ha sido aprobado y no tiene interés")
		} else {
			fmt.Println("Su prestamo ha sido aprobado y tiene interés")
		}
	} else {
		fmt.Println("Su prestamo NO ha sido aprobado")
	}
}
