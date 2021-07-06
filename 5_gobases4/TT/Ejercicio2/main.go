package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type Cliente struct {
	Legajo         int
	NombreApellido string
	DNI            int
	Telefono       int
	Domicilio      string
}

func aumentarLejago(numero int) (leg int) {
	leg = numero + 1
	return
}

func (c *Cliente) crearCliente(leg int, nom string, dni int, tel int, dom string) {
	c.Legajo = leg
	c.NombreApellido = nom
	c.DNI = dni
	c.Telefono = tel
	c.Domicilio = dom
}

func existeCliente(c Cliente, archivo string) (existe bool) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	content, err := ioutil.ReadFile(archivo)
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Print(content)
	}
	return
}

func validarCliente(c Cliente) (correcto bool, e error) {
	switch {
	case c.NombreApellido == "":
		e = errors.New("nombre vacío")
	case c.DNI == 0:
		e = errors.New("DNI vacío")
	case c.Telefono == 0:
		e = errors.New("teléfono vacío")
	case c.Domicilio == "":
		e = errors.New("domicilio vacío")
	case c.Legajo == 0:
		e = errors.New("legajo vacío")
	default:
		correcto = true
	}
	return
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores")
		fmt.Println("No quedaron archivos abiertos")
	}()

	var legajo int = 1
	c1 := Cliente{}
	c1.crearCliente(legajo, "Juan Perez", 35852586, 47025845, "Larsen")
	legajo = aumentarLejago(legajo)
	c2 := Cliente{}
	c2.crearCliente(legajo, "Pedro Gomez", 2589654, 4785412, "Zufriategui")
	legajo = aumentarLejago(legajo)
	c3 := Cliente{
		Legajo:         legajo,
		NombreApellido: "Pepe",
	}

	existeCliente(c1, "customers.txt")

	correcto, err := validarCliente((c1))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Cliente válido", correcto)
	}

	_, err1 := validarCliente((c3))
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println("Cliente válido", correcto)
	}

}
