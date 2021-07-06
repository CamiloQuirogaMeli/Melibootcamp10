package main

import (
	"fmt"
)

func main() {
	var numberMonth int
	var months = map[int]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Print("Enter number to know which month corresponds to: ")
	fmt.Scanln(&numberMonth)
	fmt.Println(numberMonth,",", months[numberMonth]);
	/*
		OTRA OPCION ES UN SWITCH CON TODOS LOS CASE DEL VALOR DE CADA MES ES DECIR CASE DEL 1 AL 12
		ELEGI ESTA POR QUE LA ENCONTRE RAPIDA
	*/
}