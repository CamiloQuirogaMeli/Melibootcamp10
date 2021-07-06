package main

import (
	"fmt"
	"io/ioutil"
)

type miError struct {
	msg string
}

func (e *miError) Error() string {
	return fmt.Sprintf("error: %s", e.msg)
}

type Cliente struct {
	legajo    int
	nombre    string
	apellido  string
	DNI       int
	telefono  int
	domicilio string
}

func (clien *Cliente) cargoCliente() {
	clien.nombre = "Franco"
	clien.apellido = "Ballet"
	clien.DNI = -8
	clien.telefono = 99371635
	clien.domicilio = "Dr Juan A Rodriguez 1415"
}

func comprobarCliente() {
	defer func() {
		errs := recover()
		if errs != nil {
			fmt.Println(errs)
		}
	}()

	fmt.Println("Comprobando que exite cliente en la base de datos")
	_, err := ioutil.ReadFile("./customers.txt")
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
}

func comprobarCampos(cliente *Cliente) (int, error) {
	if cliente.DNI <= 0 {
		err := &miError{
			msg: "el numero de DNI no puede ser inferior a 1",
		}
		return 0, err
	}
	if cliente.telefono <= 0 {
		err := &miError{
			msg: "el numero de telefono no puede ser inferior a 1",
		}
		return 0, err
	}
	return 1, nil
}

func main() {
	cli := &Cliente{
		legajo: 2888,
	}
	cli.cargoCliente()
	comprobarCliente()
	_, err := comprobarCampos(cli)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fin de la ejecucion")
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println("No han quedado archivos abiertos")

}
