package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

/*
Ejercicio 2 - Registrando clientes
El mismo estudio del ejercicio anterior, también solicita una funcionalidad para poder
registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
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
ejercicio anterior). Claramente, este archivo que intentas leer no existirá, por lo que la
función a utilizar, para intentar la lectura, devolverá un error que deberás manipular
adecuadamente como hemos visto en el contenido hasta aquí. El error deberá generar un panic,
luego lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está
dañado”, y continuar con la ejecución del programa normalmente.
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

/*Estructura cliente*/
type cliente struct {
	numLegajo   int
	nombre      string
	apellido    string
	DNI         int
	numTelefono int
	domicilio   string
}

/*Asignación previa del num de legajo*/
func (c *cliente) asignarLegajo(numero int) {

	c.numLegajo = numero

}

/*Asignación de datos*/
func (c *cliente) registrarCliente(nomb, ape, dir string, dni, telef int) {

	c.nombre = nomb
	c.apellido = ape
	c.DNI = dni
	c.numTelefono = telef
	c.domicilio = dir

}

/*Verificación de si el cliente está registrado*/
func verificaCliente(filename string, c cliente) {

	_, err := ioutil.ReadFile(filename)

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("!!!! PANIC, el archivo indicado no fue encontrado o está dañado")
	}

}

/*Validación de campos*/
func validarCampos(c cliente) (bool, error) {

	switch {
	case c.numLegajo == 0:
		return false, errors.New("error: el campo legajo se encuentra vacío")
	case c.nombre == "":
		return false, errors.New("error: el campo nombre se encuentra vacío")
	case c.apellido == "":
		return false, errors.New("error: el campo apellido se encuentra vacío")
	case c.DNI == 0:
		return false, errors.New("error: el campo DNI se encuentra vacío")
	case c.numTelefono == 0:
		return false, errors.New("error: el campo telefono se encuentra vacío")
	case c.domicilio == "":
		return false, errors.New("error: el campo domicilio se encuentra vacío")
	}

	return true, nil

}

func main() {

	var c1 cliente
	c1.asignarLegajo(3421)
	c1.registrarCliente("", "apellido", "direccion", 12345678, 987654)

	filename := "Customers.txt"
	verificaCliente(filename, c1)

	_, err := validarCampos(c1)

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")

	}()

	if err != nil {
		panic("!!!! PANIC, los datos ingresados son inválidos")
	}

}
