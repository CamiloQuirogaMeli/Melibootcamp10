package main

import (
	"fmt"
)

type Matrix struct {
	height int
	width  int
}

func (m Matrix) print() {

	for i := 0; i < m.height; i++ {
		var chain string
		for j := 0; j < m.width; j++ {
			chain += "* \t"
		}
		fmt.Println(chain)
	}
}
func (m *Matrix) set() {
	var height, width int
	fmt.Println("Ingrese la altura de la matrix")
	fmt.Scan(&height)
	fmt.Println("Ingrese el ancho de la matrix")
	fmt.Scan(&width)

	m.height = height
	m.width = width
}

func main() {
	matrixInit := Matrix{4, 4}
	var option int
	var aux int = 1

	const banner string = "\n███╗   ███╗    ███████╗    ███╗   ██╗    ██╗   ██╗\n████╗ ████║    ██╔════╝    ████╗  ██║    ██║   ██║\n██╔████╔██║    █████╗      ██╔██╗ ██║    ██║   ██║\n██║╚██╔╝██║    ██╔══╝      ██║╚██╗██║    ██║   ██║\n██║ ╚═╝ ██║    ███████╗    ██║ ╚████║    ╚██████╔╝\n╚═╝     ╚═╝    ╚══════╝    ╚═╝  ╚═══╝     ╚═════╝ "

	const options string = "1. Ver Matrix \n2. Editar Matrix \n0. Salir"

	for aux != 0 {
		fmt.Println(banner)
		fmt.Println(options)
		fmt.Scan(&option)

		switch option {
		case 1:
			matrixInit.print()
		case 2:
			matrixInit.set()

		case 0:
			aux = 0
		default:
			fmt.Println("Opcion incorrecta!!!")
		}

	}

}
