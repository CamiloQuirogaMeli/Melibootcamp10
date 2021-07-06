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
	aQueMesCorresponde()
}

func aQueMesCorresponde() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("A qué mes corresponde?")
	fmt.Println("Input format: monthNumber, ex:")
	fmt.Println("-> 7 -> Julio")
	fmt.Println("---------------------")
  
	meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"} 

	for {
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  // convert CRLF to LF (Only Windows)
	  text = strings.Replace(text, "\n", "", -1)
	  num, _ := strconv.ParseInt(text, 10, 64)
	  num--
	  if num < 1 || num > 12 {
		fmt.Println("No existe tal mes")
	  } else { fmt.Println("El numero", num, "corresponder al mes", meses[num]) }
	}
}