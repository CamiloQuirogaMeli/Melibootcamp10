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
por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y
continuar con la ejecución del programa normalmente.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Persona struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       string
	Telefono  string
	Domicilio string
}

func main() {
	id := rand.Intn(300)
	persona := &Persona{Legajo: id}

	customer, err := ioutil.ReadFile("customers.txt")
	if err != nil {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("No se pudo conectar con el archivo")
				fmt.Println("ejecución finalizada")
			}

		}()
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {

		if customer[0] != byte(id) {
			fmt.Println("El usuario ya se encuentra registrado")
		} else {
			fmt.Println(persona)
			fmt.Println("El usuario ha sido registrado correctamente")
		}
	}
}
