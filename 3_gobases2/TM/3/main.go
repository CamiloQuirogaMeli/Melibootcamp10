package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

/* Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que se le pueda pasar como parámetros la cantidad de
minutos trabajados por mes y la categoría, y nos devuelva su salario */

func main() {
	calcularSalario()
}

func calcularSalario() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calcular salario")
	fmt.Println("Minutos por mes + categoria ej: 1000 C")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF (Only Windows)
		text = strings.Replace(text, "\n", "", -1)
		arr := strings.Fields(text)
		if (len(arr) < 1) {
			fmt.Println("No escribiste nada :C")
			continue
		}
		horas, _ := strconv.ParseFloat(arr[0], 64)
		salario := 0.0
		horas = horas / 60

		switch strings.ToLower(arr[1]) {
			case "a": {
				salario = horas * 3000
			}
			case "b": {
				salario = horas * 1500
			}
			default: {
				salario = horas * 1000
			}
		}
		fmt.Printf("Salario: %0.2f\n", salario)
	}
}