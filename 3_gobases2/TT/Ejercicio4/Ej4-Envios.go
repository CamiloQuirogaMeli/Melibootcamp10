package main

import "fmt"

type productos struct{
	tamanio float64
	tipoProducto string
}

func tamTotal (tipoProd string, tam float64)(float64, string){
	switch tipoProd {
	case "chico":
		return tam, tipoProd
	case "mediano":
		return tam * 1.05, tipoProd
	case "grande":
		return tam * 1.20, tipoProd
	case "fragil":
		return tam * 1.75, tipoProd
	case "especial":
		return tam, tipoProd
	}
	return 0, ""
}

func agregarProd(tipoProd string, tamProd float64, dispEnvios float64, prodSlice []productos) ([]productos, float64){

	if dispEnvios > 0{

		dispEnvios = dispEnvios - tamProd
		prodSlice = append (prodSlice, productos{tipoProducto: tipoProd, tamanio: float64(tamProd)})

	}

	return prodSlice, dispEnvios
}

func main(){
	var prepProd []productos //lista prod a enviar
	var dispEnvios float64 //cant de tamaño disponible para hacer envios
	pEspeciales := false

	prod1 := productos{tamanio: 500000, tipoProducto: "chico"}
	prod2 := productos{tamanio: 1000000, tipoProducto: "mediano"}
	prod3 := productos{tamanio: 1500000, tipoProducto: "grande"}
	prod4 := productos{tamanio: 400000, tipoProducto: "fragil"}
	prod5 := productos{tamanio: 250000, tipoProducto: "especial"}

	totDimension, tipoProd := tamTotal(prod1.tipoProducto, prod1.tamanio)
	prepProd, dispEnvios = agregarProd(prod1.tipoProducto, prod1.tamanio, dispEnvios, prepProd)
	totDimension, tipoProd = tamTotal(prod2.tipoProducto, prod2.tamanio)
	prepProd, dispEnvios = agregarProd(prod2.tipoProducto, prod2.tamanio, dispEnvios, prepProd)
	totDimension, tipoProd = tamTotal(prod3.tipoProducto, prod3.tamanio)
	prepProd, dispEnvios = agregarProd(prod3.tipoProducto, prod3.tamanio, dispEnvios, prepProd)
	totDimension, tipoProd = tamTotal(prod4.tipoProducto, prod4.tamanio)
	prepProd, dispEnvios = agregarProd(prod4.tipoProducto, prod4.tamanio, dispEnvios, prepProd)

	for _, prods := range prepProd{
		if prods.tipoProducto != "especial"{
			pEspeciales = true
		}
	}
	if pEspeciales != true{
		totDimension, tipoProd = tamTotal(prod5.tipoProducto, prod5.tamanio)
		prepProd, dispEnvios = agregarProd(prod5.tipoProducto, prod5.tamanio, dispEnvios, prepProd)
	}else{
		fmt.Println("Hay productos no especiales en el envio")
	}
	fmt.Println("Pedidos cargados: ", prepProd)
	fmt.Println("Disponibilidad restante ", dispEnvios)
}


/*Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un adicional al momento de ser enviado:

Chico: Ningún adicional.
Mediano: Requiere un %5 más de espacio
Grande: Requiere un %20 más de espacio
Frágil: Requiere un %75 más de espacio
Especial: Sólo puede ser enviado con productos especiales

Se solicita que:
Los productos guarden el tamaño y tengan un método Tamaño Total que nos devuelva el espacio en cm3 que requerimos para ser enviado.
Y una estructura Flete que tenga los métodos:
Agregar Producto: agregar producto al flete.
Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede cargar un total de 10.000.000 cm3 por envío.
*/
