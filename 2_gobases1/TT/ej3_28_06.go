package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var j int
	for {
		fmt.Println("Ingrese su edad")
		fmt.Println("---------------------")
		fmt.Print("-> ")
		var i int
		fmt.Scan(&i)
		if i < 22 {
			fmt.Println("Los prestamos son solo para mayores de 22 años, disculpe las molestias!")
			break
		}

		fmt.Println("En la actualidad se encuentra empleado?")
		fmt.Println("Ingrese SI/NO")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text != "SI" {
			fmt.Println("Los prestamos son solo para personas empleadas, disculpe las molestias!")
			break
		}

		fmt.Println("Tiene una antiguedad mayor a 1 año?")
		fmt.Println("Ingrese SI/NO")
		textA, _ := reader.ReadString('\n')
		textA = strings.Replace(textA, "\n", "", -1)

		if textA != "SI" {
			fmt.Println("Los prestamos son solo para personas con mas de 1 año de antiguedad, disculpe las molestias!")
			fmt.Scan(&j)
			break
		}

		fmt.Println("Usted es apto para obtener un prestamo")

		fmt.Println("Ingrese su salario")
		var k int
		fmt.Scan(&k)
		if k > 100000 {
			fmt.Println("Su credito no tendra intereses")
		}
		break
	}

}
