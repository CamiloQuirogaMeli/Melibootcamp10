package main

import f "fmt"

func main() {

	/*
	   Ejercicio 4 - A qué mes corresponde

	   Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
	   que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
	   por qué?
	   Ej: 7, Julio

	   Respuesta:
	   Con un map, siendo el numero del mes la clave.
	   Con un switch o un IF, tambien se puede resolver.
	   Entiendo que con un package de time, se podria resolver facilmente. 
	   Ahora eligo el map, ya que no tengo conocimiento total del package time. 
	*/
	var numMonth uint
	f.Println("Escriba un numero de mes")
	f.Scan(&numMonth)
	var monthMap = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8:"Agosto", 9: "Septiembre", 10:"Octubre", 11: "Noviembre", 12:"Diciembre"}

	for key, element := range monthMap {
		if key == int(numMonth) {
			f.Println(key,",",element)
		}
	}

//Forma dos, Switch
	// var nameMonth string

	// switch numMonth {
	// case 1:
	// 	nameMonth = "Enero"
	// case 2:
	// 	nameMonth = "Febrero"
	// case 3:
	// 	nameMonth = "Marzo"
	// case 4:
	// 	nameMonth = "Abril"
	// case 5:
	// 	nameMonth = "Mayo"
	// case 6:
	// 	nameMonth = "Junio"
	// case 7:
	// 	nameMonth = "Julio"
	// case 8:
	// 	nameMonth = "Agosto"
	// case 9:
	// 	nameMonth = "Septiembre"
	// case 10:
	// 	nameMonth = "Octubre"
	// case 11:
	// 	nameMonth = "Noviembre"

	// case 12:
	// 	nameMonth = "Diciembre"
	// default:
	// 	nameMonth = "null"
	// }


	// if nameMonth != "null" {
	// 	f.Println(numMonth, ", ", nameMonth)
	// }else{
	// 	println("No exisite ese mes")
	// }
}
