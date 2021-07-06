package main

import (
	"fmt"
	"os"
)

func main() {
	ejercicio2()
}

func ejercicio2() {
	texto, err := os.ReadFile("../Ejercicio1/ejercicio1")
	fmt.Printf("ID\tPRECIO\tCANTIDAD\n")
	textoCompleto := string(texto)
	if err == nil {
		//productoDividido := strings.Split(string(textoCompleto), ";")
		for i := range textoCompleto {
			if string(textoCompleto[i]) != "\n" {
				if string(textoCompleto[i]) != ";" {
					fmt.Print(string(textoCompleto[i]))
				} else {
					fmt.Print("\t")
				}
			} else {
				fmt.Println()
			}

		}
	} else {
		fmt.Println(err)
	}

}
