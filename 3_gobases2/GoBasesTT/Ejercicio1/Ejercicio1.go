package main

import "fmt"

type Estudiante struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

//metodo detalle
func (e Estudiante) metMostrarEstudiante() {
	fmt.Println("Nombre: " + e.nombre)
	fmt.Println("Apellido: " + e.apellido)
	fmt.Println("Dni: " + e.dni)
	fmt.Println("Fecha: " + e.fecha)
	fmt.Println("")
}

//metodo detalle
func (e Estudiante) metCargarEstudiante() Estudiante {
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s", &e.nombre)
	fmt.Println("Ingrese su apellido: ")
	fmt.Scanf("%s", &e.apellido)
	fmt.Println("Ingrese su dni: ")
	fmt.Scanf("%s", &e.dni)
	fmt.Println("Ingrese su fecha: ")
	fmt.Scanf("%s", &e.fecha)

	return e
}

//imprimir los datos mediante un metodo
func main() {

	//est := Estudiante{nombre: "Santiago", apellido: "Esposito", dni: "34206486", fecha: "19/04/1989"}
	est := Estudiante{}
	est1 := est.metCargarEstudiante()
	est1.metMostrarEstudiante()
}
