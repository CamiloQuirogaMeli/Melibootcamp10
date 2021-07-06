package main

import "fmt"

func main() {
	var (
		age                 int     = 23
		hasWork             bool    = true
		salary              float64 = 60000
		months_of_seniority int     = 13
	)

	if age > 22 && hasWork && months_of_seniority > 12 {
		if salary > 100000 {
			fmt.Println("Cumple con todos los requisitos, adquiera su credito ya y ademas, NO LE COBRAREMOS INTERESES")
		} else {
			fmt.Println("Cumple con todos los requisitos, adquiera su credito ya con tasas de interes preferenciales")
		}

	} else {
		fmt.Println("Lo sentimos, no cumple los requisitos para poder darle un prestamo")
	}
}
