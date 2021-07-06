package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

/* Ejercicio 2 - Calcular promedio
Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo */

func main() {
	calcularPromedio()
}

func calcularPromedio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calcular promedio")
	fmt.Println("Notas separadas por espacios, ej: 10 10 10")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF (Only Windows)
		text = strings.Replace(text, "\n", "", -1)
		arr := strings.Fields(text)
		
		res, err := withEllipsis(arr...)
		if err != nil {
			fmt.Println("Error. No puedes poner numeros negativos.")
		} else {
			fmt.Println("Promedio:", res)
		}
	}
}

func withEllipsis(calificaciones ...string) (int, error) {
	promedio, c := 0, 0
	for _, val := range calificaciones {
		val2, _ := strconv.Atoi(val)
		if val2 < 0 {
			return 0, errors.New("No se pueden poner numeros negativos.")
		} else {
			promedio += val2
			c++
		}
	}
	return promedio / c, nil
}