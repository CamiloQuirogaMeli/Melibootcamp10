package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

/* Ejercicio 1 - Impuestos de salario

Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
depositar el sueldo, para cumplir el objetivo necesitan una función que nos devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un %17 del
sueldo y si gana más de $150.000 se le descontará además un %10. */

func main() {
	impuestoDeSalario()
}

func impuestoDeSalario() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Impuesto de salario")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF (Only Windows)
		text = strings.Replace(text, "\n", "", -1)
		salario, _ := strconv.Atoi(text)
		if salario < 50001 {
			fmt.Println("Salarios menores a $50.001 no pagan impuestos.")
		} else if salario < 150001 {
			fmt.Println("Salarios mayores a $50.000 pagan impuestos del %17 del sueldo.")
			fmt.Println("Impuesto total en pesos:", salario * 17 / 100)
			fmt.Println("Salario con el impuesto aplicado:", salario - (salario * 17 / 100))
		} else {
			fmt.Println("Salarios mayores a $150.000 pagan impuestos del %27 del sueldo.")
			fmt.Println("Impuesto total en pesos:", salario * 27 / 100)
			fmt.Println("Salario con el impuesto aplicado:", salario - (salario * 27 / 100))
		}
	}
}