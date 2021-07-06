package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	legajo    string
	nombre    string
	apellido  string
	dni       int
	telefono  int
	domicilio string
}

func (c *Cliente) corroborarDatos() (string, error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error en los datos ingresados")
		}
	}()

	if c.legajo != "" && c.nombre != "" && c.apellido != "" && c.dni > 0 && c.telefono > 0 && c.domicilio != "" {
		fmt.Println(c.legajo)
		return "Datos del cliente completos", nil
	} else {
		return "", errors.New("Datos incompletos")
	}
}

func leerArchivoClientes(nombreArchivo string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo indicado no fue encontrado o está dañado")
			fmt.Println("No han quedado archivos abiertos")

		}
	}()

	fmt.Println("Se intentará leer el archivo")

	dat, err := os.Open(nombreArchivo)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(dat)
	}
}

func main() {
	/*

		El mismo estudio del ejercicio anterior, también solicita una funcionalidad para poder registrar
		datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
		- Legajo
		- Nombre y Apellido
		- DNI
		- Número de teléfono
		- Domicilio
		El número de legajo debe ser asignado o generado por separado y en forma previa a la carga
		de los restantes gastos.
		Para cumplir con los objetivos de aprendizaje de este módulo, aún no vamos a escribir los
		datos en un archivo. Sin perjuicio de lo cual, supongamos que quieres constatar si el cliente
		se encuentra ya registrado y para ello precisas leer los datos de un archivo .txt. En algún
		lugar de tu código, simula que intentas leer un archivo llamado “customers.txt” (como en el
		ejercicio anterior). Claramente, este archivo que intentas leer no existirá, por lo que la función
		a utilizar, para intentar la lectura, devolverá un error que deberás manipular adecuadamente
		como hemos visto en el contenido hasta aquí. El error deberá generar un panic, luego lanzar
		por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.
		Todos los campos son requeridos. De modo que, luego de intentar constatar la existencia del
		cliente a registrar, desarrolla una función para validar que todos los campos contienen un
		valor distinto de cero, retornando al menos dos valores. Uno de los valores retornados
		deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero
		(recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).
		Utiliza recover para recuperar el valor de los panics que puedan surgir.
		Antes de finalizar la ejecución, incluso si surgen panics, se deberá imprimir por consola los
		siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de
		ejecución” y “No han quedado archivos abiertos” (en ese orden).
		No olvides realizar las validaciones necesarias para cada retorno que pueda contener un
		valor error (por ejemplo las que intenten leer archivos). Genera algún error, personalizandolo
		a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la
		validación pertinente para el caso de error retornado).


	*/
	nombreArchivo := "customers.txt"
	salir := false
	var opcion int
	cliente := Cliente{}

	for !salir {
		fmt.Println("Digita una opción:")
		fmt.Println("\t 1: Ingresar un cliente y corroborar la información")
		fmt.Println("\t 2: Escribir los datos el archivo customers.txt")
		fmt.Println("\t 3: Salir")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:

			fmt.Println("Digita el legajo del cliente")
			fmt.Scanln(&cliente.legajo)
			fmt.Println("Digita el nombre del cliente")
			fmt.Scanln(&cliente.nombre)
			fmt.Println("Digita el apellido del cliente")
			fmt.Scanln(&cliente.apellido)
			fmt.Println("Digita el dni del cliente")
			fmt.Scanln(&cliente.dni)
			fmt.Println("Digita el telefono del cliente")
			fmt.Scanln(&cliente.telefono)
			fmt.Println("Digita el domicilio del cliente")
			fmt.Scanln(&cliente.domicilio)

			status, err := cliente.corroborarDatos()
			if err != nil {
				fmt.Println("Se detectaron varios errores en tiempo de ejecución")
			} else {
				fmt.Println(status)
			}

		case 2:
			leerArchivoClientes(nombreArchivo)

		case 3:
			salir = true
		}
		fmt.Println("Ejecución finalizada")

	}
}
