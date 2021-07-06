package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Cliente struct {
	legajo         int
	nombreApellido string
	dni            int
	telefono       int
	domicilio      string
}

func (c *Cliente) asignarLegajo(legajo int) (bool, error) {
	if legajo == 0 {
		return false, errors.New("el legajo no puede ser cero")
	}

	c.legajo = legajo

	return true, nil
}

func (c *Cliente) asignarDatosRestantes(nombre, domicilio string, dni, telefono int) (bool, error) {
	if nombre == "" || domicilio == "" {
		return false, errors.New("el nombre y domicilio no pueden ser cadenas vacias")
	}

	if dni == 0 || telefono == 0 {
		return false, errors.New("el dni y el telefono no pueden ser cero")
	}

	c.nombreApellido = nombre
	c.dni = dni
	c.telefono = telefono
	c.domicilio = domicilio

	return true, nil
}

func yaExiste(data string) (bool, error) {
	return false, nil
}

func imprimirMensaje1() {
	fmt.Println("No han quedado archivos abiertos")
}

func imprimirMensaje2() {
	fmt.Println("Se detectaron varios errores en tiempo de ejecuión")
}

func imprimirMensaje3() {
	fmt.Println("Fin de la Ejecución")
}

func main() {
	fmt.Println("ejecución inicializada")

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	defer imprimirMensaje1()
	defer imprimirMensaje2()
	defer imprimirMensaje3()

	// leo los datos del archivo
	data, err := ioutil.ReadFile("./customers.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

	// valido si el cliente ya existe
	yaExiste, er := yaExiste(string(data))

	if er != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

	cliente := Cliente{}

	if !yaExiste {
		// valido que los datos esten completos
		_, errLegajo := cliente.asignarLegajo(3)

		if errLegajo != nil {
			panic(errLegajo)
		}

		// valido el resto de los datos
		_, errDatosRestantes := cliente.asignarDatosRestantes("Mario Rosales", "Santa Fe", 121212, 343434455)

		if errDatosRestantes != nil {
			panic(errDatosRestantes)
		}
	}
}
