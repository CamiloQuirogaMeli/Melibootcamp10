package main

import (
	"fmt"
)

func main() {
	// variable nombre bien definida
	var nombre string
	fmt.Println("nombre: ", nombre)
	// variable apellido bien definida
	var apellido string

	// esta variable esta mal definida
	// la definicion correcta seria invirtiendo el nombre y el tipo
	// var int edad
	//forma correcta es var edad int
	var edad int
	fmt.Println("edad: ", edad)

	// la sintaxis es correcta, pero no es correcto el tipo de dato
	// la variable apellido se declaro como tipo string y se le esta queriendo volver a difinir como tipo int y asignar un numero
	// quizas pueda asiganrse el valor 6 pero como string
	// apellido := 6
	// forma correcta
	apellido = "6"
	fmt.Println("apellido: ", apellido)

	// aca falta aclarar el tipo de dato, que en este caso deberia ser bool, por el tipo de dato que se le esta asignando
	// var licencia_de_conducir = true
	var licencia_de_conducir bool = true
	fmt.Println("licencia_de_conducir: ", licencia_de_conducir)

	// esto posee error de sintaxis ya que las variables no pueden ser definidas con espacios
	// Las palabras deben ser concatenadas y las primeras letras de cada palabra va en mayusculas
	// var estatura de la persona int
	// La forma correcta es
	var estaturaDeLaPersona int
	fmt.Println("estaturaDeLaPersona: ", estaturaDeLaPersona)

	// esta variable se esta definiendo e inicializando de forma correcta
	cantidadDeHijos := 2
	fmt.Println("cantidadDeHijos: ", cantidadDeHijos)

	/**
	Cabe destacar que las variables, si son definidas deben ser utilizadas.
	De otra manera el programa no va a compilar
	*/
}
