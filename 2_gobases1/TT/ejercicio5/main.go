package main

import "fmt"

func main()  {
var alumnos = []string{"Benjamín", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel",
						"Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	aniadir :=
		`
	Seleccione:
	[ 1 ] AnadirAlumno
	[ 2 ] NoAnadirAlumno
	
`
	fmt.Print(aniadir)

	var eleccion int
	fmt.Scanln(&eleccion)


	switch eleccion{
	case 1:
		nombre:= `Ingrese nombre`
		fmt.Print(nombre)

		var name string
		fmt.Scan(&name)

		alumnos = append(alumnos, name)

		fmt.Println(alumnos)

	default:
		fmt.Println(alumnos)
	}


}


