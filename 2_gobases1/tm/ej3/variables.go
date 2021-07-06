package ej3

import "fmt"

func Variables()  {
	exam := [7][2]string{								//Correct
		{"var 1nombre string","false"},				//var nombre string
		{"var apellido string","true"},
		{"var int edad","false"},					//var edad int
		{"1apellido := 6","false"},					//apellido := 6
		{"var licencia_de_conducir = true","true"},
		{"var estatura de la persona int","false"}, //var estaturaDeLaPersona int
		{"cantidadDeHijos := 2","true"},
	}

	for i := 0; i < 7; i++ {
		fmt.Print(exam[i][0])
		if exam[i][1] == "false"{
			fmt.Print(" ----> Wrong!!!\n")
		}else{
			fmt.Print(" <-------------> Correct\n")
		}
	}
}