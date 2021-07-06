package main
import (
	"fmt"
)

/* Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda a imprimir la edad de Benjamin. Por otro lado también es necesario:
- Saber cuántos de sus empleados son mayores a 21 años.

3

- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
- Eliminar a Pedro del mapa. */

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	queEdadTiene()
}

func queEdadTiene() {
	fmt.Println(employees["Pedro"])
	employees["Federico"] = 25
	delete(employees, "Pedro")
	//fmt.Println(employees["Federico"])
}