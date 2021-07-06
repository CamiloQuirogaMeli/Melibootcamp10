package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func salario(minutos float64, categoria string) float64 {
	var resultado float64
	var horas float64 = minutos / 60
	if categoria == "a" {
		resultado = 3000 * horas * 1.5
	}
	if categoria == "b" {
		resultado = 1500 * horas * 1.2
	}
	if categoria == "c" {
		resultado = 1000 * horas
	}
	return resultado
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese los minutos trabajados")
	var i float64
	fmt.Scan(&i)
	fmt.Println("Ingrese la categoria que es")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	a := salario(i, text)

	fmt.Println("Su salario es:", a)

}
