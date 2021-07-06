package main

import "fmt"

func main() {

	var MONTHS_MAP = map[int]string{1: "January", 2: "Februry", 3: "April", 4: "May", 5: "June", 6: "July", 7: "August", 8: "September", 9: "October", 10: "November", 12: "December"}
	var month int

	fmt.Println("Entry month number")
	fmt.Scanln(&month)

	fmt.Printf("The month [%d] is: [%s]\n", month, MONTHS_MAP[month])

	/*
		Debido a que son constantes, se podria realizar un array con los 12 meses del año y
		simplemente obtener el valor del mismo a traves del mes que nos indiquen. Seria O(1) por ende seria eficiente.

		Elegi usar un MAP porque es O(1) y me parecio que era lo indicado.
	*/

}
