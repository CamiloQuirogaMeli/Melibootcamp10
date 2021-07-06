package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var option int
	var aux int = 1

	for aux != 0 {

		const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

		const options string = "1. Leer archivo customers.txt \n0. Salir"

		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:

			watchFile()
			fmt.Println("Ejecucion finalizada")

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}
	}
}

func watchFile() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	dat, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}

}
