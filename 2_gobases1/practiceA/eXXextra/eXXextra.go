// Ejercicio 1 - Letras de una palabra

// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
// una de las letras por separado para deletrearla, para ello es necesario crear una aplicación
// tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego
// imprimí cada una de las letras

package main

import (
	"fmt"
	"strings"
)

func main() {
	var option int
	words := make(map[string]int)
	for {
		for {
			fmt.Printf("\nIngrese una opción:\n1- Agregar una palabra. \n2- Buscar una palabra. \n3- Listar palabras. \n4- Deletrear. \n5- Salir. \n-> ")
			if _, err := fmt.Scanln(&option); err != nil {
				fmt.Println("Error:", err)
				fmt.Println()
			} else {
				break
			}
		}

		switch option {
		case 1:
			{ // Guardar palabra
				var word string
				fmt.Printf("Escribe una palabra: ")
				fmt.Scanf("%s\n", &word)
				words[strings.ToLower(word)] = len(word)
			}
		case 2:
			{ // Buscar palabra
				var word string
				fmt.Printf("Buscar: ")
				fmt.Scanln("%s", &word)
				_, exists := words[strings.ToLower(word)]
				if exists != false {
					fmt.Println(strings.ToUpper(word), "tiene ", words[strings.ToLower(word)], " letras.")
				} else {
					fmt.Println("No se encontró la palabra", word)
				}
			}
		case 3:
			{ // Listar palabras
				for key, value := range words {
					fmt.Println(strings.ToUpper(key), "tiene ", value, " letras.")
				}
			}
		case 4:
			{ // Deletrear palabras
				for key, value := range words {
					fmt.Println(strings.ToUpper(key), "(", value, " letras): ")
					for i, letter := range key {
						fmt.Println(i, "- ", strings.ToUpper(string(letter)))
					}
					fmt.Println()
				}
			}
		default:
			fmt.Println("\nOpción incorecta.")
		}
		if option == 5 {
			break
		}
	}
}
