package ejercicios

import "fmt"

var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func SextoEjercicio() {
	fmt.Printf("La edad de Benjamin es %d\n", empleados["Benjamin"])
	imprimirEmpleadosMayoresDe21()
	empleados["Federico"] = 25
	fmt.Println("Hemos añadido al empleado Federico")
	imprimirEmpleadosMayoresDe21()
	delete(empleados, "Pedro")
	fmt.Println("Hemos elminado al empleado Pedro")
	imprimirEmpleadosMayoresDe21()
}

func calcularEmpleadosMayoresDe21() int {
	var empleadosMayoresDe21 = 0
	for _, empleado := range empleados {
		if empleado > 21 {
			empleadosMayoresDe21++
		}
	}
	return empleadosMayoresDe21
}

func imprimirEmpleadosMayoresDe21() {
	fmt.Printf("La cantidad de empleados mayores de 21 es: %d\n", calcularEmpleadosMayoresDe21())
}
