package main

import (
	"fmt"
	"io/ioutil"
)

type Cliente struct {
	legajo         int
	nombreCompleto string
	dni            int
	telefono       int
	domicilio      string
}

func main() {
	data := leerArchivo("./customers.txt")
	cliente := nuevoCliente()
	cargarCliente(data, cliente, "./customers.txt")
	fmt.Println("Ejecución finalizada")
}

func cargarCliente(data string, cliente Cliente, archivo string) {
	// simular carga de cliente en el archivo
	fmt.Println("Cliente añadido con éxito")
}

func nuevoCliente() Cliente {
	var cliente Cliente
	fmt.Printf("*** CARGAR NUEVO CLIENTE ***\n")
	fmt.Printf("Nombre: ")
	fmt.Scanf("%s", &cliente.nombreCompleto)
	fmt.Printf("DNI: ")
	fmt.Scanf("%d", &cliente.dni)
	fmt.Printf("Telefono: ")
	fmt.Scanf("%d", &cliente.telefono)
	fmt.Printf("Domicilio: ")
	fmt.Scanf("%s", &cliente.domicilio)
	return cliente
}

func leerArchivo(file string) string {
	dataByte, err := ioutil.ReadFile("./customers.txt")
	data := string(dataByte)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		return data
	}
}
