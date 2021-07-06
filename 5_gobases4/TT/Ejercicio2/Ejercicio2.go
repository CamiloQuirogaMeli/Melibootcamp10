package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Ejercicio 2 - Registrando clientes
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
		por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y
		continuar con la ejecución del programa normalmente.

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
	c := Cliente{333, "", "Olivera", 33334124, 333242, "Calle 001"}

	fmt.Print(c)

	defer fmt.Print("Fin de la ejecucion!")
	leerDoc()
	c.validarCliente()

}

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	NumeroTel int
	Domicilio string
}

func leerDoc() string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	datosArchivo, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("error: el archivo no fue encontrado o esta dañado")
	}
	return string(datosArchivo)
}

func (c *Cliente) validarCliente() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if c.Legajo == 0 || c.DNI == 0 || c.NumeroTel == 0 || c.Nombre == "" || c.Apellido == "" || c.Domicilio == "" {
		panic("Se detectaron varios errores en tiempo de ejecucion")
	}
}
