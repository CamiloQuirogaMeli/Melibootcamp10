// package declaration
package main

// library imports
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main function this is the fist function in run with program start
func main() {
	//Create variable for read console
	scanner := bufio.NewScanner(os.Stdin)

	//initialize variable for "while" loop
	decision := 0

	//in Golang "while" loop is replaced with "for" loop
	for decision == 0 {
		//message to user
		fmt.Println("Por favor introduzca la letra que desea deletrear:")
		//read input to user
		scanner.Scan()
		fmt.Println()
		//execute external function to split text
		separarLetras(scanner.Text())
		fmt.Println()

		//message to continue or terminate the program
		fmt.Println("Si desea continuar, digite el numero 0, de lo contrario digite el numero 1")
		scanner.Scan()

		//if conditional to terminate programm
		if scanner.Text() == "1" || scanner.Text() != "0" {
			fmt.Println("Muchas gracias por utilizar el programa")
			decision = 1
		}
		fmt.Println()
	}
}

func separarLetras(text string) {
	println("La cantidad de caracteres en la palabra es:", strings.Count(text, "")-1)
	//Function to Uppercase Text
	textCapitallized := strings.ToUpper(text)

	//Separate characters to text
	textSeparate := strings.Split(textCapitallized, "")

	//message to user
	println("El texto separado es:")

	//Print each character
	for i := range textSeparate {
		fmt.Println(textSeparate[i])
	}
}
