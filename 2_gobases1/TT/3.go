package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

/* Un Banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los
mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y
tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no
les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario generar una aplicación que tenga estas variables y que imprima un mensaje de
acuerdo a cada caso. */

func main() {
	prestamo()
}

func prestamo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Descuento")
	fmt.Println("Input format: age employed antiquity salary, ex:")
	fmt.Println("-> 22 yes 2 50000")
	fmt.Println("---------------------")
  
	for {
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  // convert CRLF to LF (Only Windows)
	  text = strings.Replace(text, "\n", "", -1)
	  arr := strings.Fields(text)

	  age, _ := strconv.Atoi(arr[0])
	  employed := strings.ToLower(arr[1])
	  antiquity, _ := strconv.Atoi(arr[2])
	  salary, _ := strconv.Atoi(arr[3])

	  if age < 22 {
		  fmt.Println("No se puede otorgar el prestamo si eres menor de 22 años.")
		} else if employed != "yes" && employed != "si" {
		  fmt.Println("No se puede otorgar el prestamo si no estás empleado.")
		} else if antiquity < 1 {
		  fmt.Println("No se puede otorgar el prestamo si tienes menos de un año de antigüedad.")
		} else {
		if salary < 100000 {
			fmt.Println("Se te cobrará intereses debido a que tu sueldo es menor a $100.000")
		} else {
			fmt.Println("No se te cobrará intereses debido a que tu sueldo es mejor a $100.000")
		}
	  }
	}
}