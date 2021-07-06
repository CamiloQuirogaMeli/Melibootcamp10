/*
Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
acuerdo a cada caso.
*/

// package declaration
package main

// library imports
import (
	"fmt"
	"os"
)

func main() {
	welcome := `----------------------------------------------------------------
----------------- Bienvenidos a su banco virtual --------------
----------------------------------------------------------------`
	fmt.Println(welcome)
	fmt.Println("Para iniciar la solicitud de prestamo, por favor lea atentamente y digite los datos solicitados. ->")
	fmt.Print("Por favor digite su edad ->")
	var age int = validation()
	if age <= 22 {
		fmt.Println("El prestamo no puede ser efectivo, ya que no cumple con el requerimiento de edad ya que tiene que ser mayor a 22 años")
		os.Exit(0)
	}
	fmt.Print("¿Usted se encuentra actualmente trabajando? Digite 1 si asi es, de lo contrario digite 0 -> ")
	var employee int = validation()
	if employee != 1 {
		fmt.Println("El prestamo no puede ser efectivo, ya que no cumple con el requerimiento de trabajo, para otorgarle el prestamo es importante que usted se encuentre trabajando")
		os.Exit(0)
	}
	fmt.Print("¿Cuantos meses lleva trabajando? ->")
	var monthsJob int = validation()
	if monthsJob <= 12 {
		fmt.Println("El prestamo no puede ser efectivo, ya que no cumple con el requerimiento de años de empleo, para otorgarle el prestamo es importante que usted lleve mas de 12 meses trabajando")
		os.Exit(0)
	}
	fmt.Print("¿Cuanto es su sueldo en DOLARES(USD)?, digitelo sin puntos(.) ni comas(,) ->")
	var salary int = validation()
	fmt.Println("")
	fmt.Print("El prestamo es efectivo, ")
	if salary > 100000 {
		fmt.Print(" en su caso no se le cobrara la tasa de interes")
	} else {
		fmt.Print(" en su caso se le cobrara tasa de interes")
	}

}

func validation() int {
	var value int
	_, err := fmt.Scanf("%d", &value)

	if err != nil {
		fmt.Println("Bad Input")
		os.Exit(0)
	}
	return value
}
