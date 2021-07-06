package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

//
// Ejercicio 1 - Guardar archivo
// Una empresa que se encarga de vender productos de limpieza necesita:
// 1. Implementar una funcionalidad para guardar un archivo de texto, con la información
//    de productos comprados, separados por punto y coma.
// 2. Debe tener el id del producto, precio y la cantidad.
// 3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
//

const (
	PATH = "../"
	FILE = "compras.txt"
)

type Producto struct {
	Id     int
	Precio float64
	Stock  int
}

func (p Producto) guardarArchivo(cantidad int) error {

	// Conservamos los datos anteriores del archivo
	contentTemp, _ := ioutil.ReadFile(PATH + FILE)

	valorTotal := float64(cantidad) * p.Precio
	// Datos Temp; ID; Stock; Precio(Unidad); Cantidad; Valor total
	contenido := fmt.Sprintf("%s%d;%d;%.2f;%d;%.2f\n", contentTemp, p.Id, p.Stock, p.Precio, cantidad, valorTotal)
	msg := []byte(contenido)
	err := ioutil.WriteFile(PATH+FILE, msg, 0644)

	return err
}

func (p *Producto) Comprar(cantidad int) error {
	if p.Stock > 0 {
		if cantidad > 0 && cantidad <= p.Stock {
			rsp := p.guardarArchivo(cantidad)
			if rsp != nil {
				msgError := fmt.Sprint("\t- Ocurrio un error al guardar los datos para el producto Id:", p.Id, "\n", rsp)
				return errors.New(msgError)
			} else {
				p.Stock = p.Stock - cantidad
				return nil
			}
		} else {
			msgError := fmt.Sprintf("\t- Lo sentimos debe comprar una cantidad mayor a 0 o menor a %d, para el producto Id: %d", p.Stock, p.Id)
			return errors.New(msgError)
		}
	}
	msgError := fmt.Sprint("\t- Lo sentimos no queda inventario para el producto Id:", p.Id)
	return errors.New(msgError)
}

func main() {
	// Inicializamos los productos
	p1 := Producto{
		Id:     1,
		Precio: 10000,
		Stock:  20}

	p2 := Producto{
		Id:     2,
		Precio: 5200,
		Stock:  36}

	var cantidadCompraProductos int

	fmt.Println("======= Guardar archivo =======")

	// Realizamos la compra de los productos
	cantidadCompraProductos = 5
	fmt.Printf("Compramos %d articulos del producto ID: %d\n", cantidadCompraProductos, p1.Id)
	err1 := p1.Comprar(cantidadCompraProductos)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("\t- Su compra fue exitosa para el producto ID:", p1.Id)
	}
	cantidadCompraProductos = 2
	fmt.Printf("Compramos %d articulos del producto ID: %d\n", cantidadCompraProductos, p2.Id)
	err2 := p2.Comprar(cantidadCompraProductos)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("\t- Su compra fue exitosa para el producto ID:", p2.Id)
	}
	fmt.Println("================================")
	// Mostramos el contenido del archivo
	fmt.Println("Contenido del archivo: ", PATH+FILE)
	content, err := ioutil.ReadFile(PATH + FILE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", content)
	}

}
