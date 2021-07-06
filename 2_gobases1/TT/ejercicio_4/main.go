package main

import "fmt"

func main() {
	
	var currentMonth = 10

// version basica con ifs
	fmt.Printf("El mes numero %d es %s\n", currentMonth, monthNameForNumberV1(currentMonth))

// version utilizando switch
	fmt.Printf("El mes numero %d es %s\n", currentMonth, monthNameForNumberV2(currentMonth))

// version utilizando un array
	fmt.Printf("El mes numero %d es %s\n", currentMonth, monthNameForNumberV3(currentMonth))

// version utilizando un map
	fmt.Printf("El mes numero %d es %s\n", currentMonth, monthNameForNumberV4(currentMonth))

// en lo personal, prefiero la alternativa de map porque se pueden modificar los valores y se adapta en todos los
// lugares que se esta utilizando el mismo, a diferencia del if y switch
// otro beneficio que tiene es lo que se utiliza para indexar, ya que puede ser cualquier tipo y se puede adaptar
// el map facilmente a este, a diferencia del array que se necesitan numeros enteros para poder indexar los distintos valores
}

func monthNameForNumberV1(currentMonth int) string {

	var monthName string

	if currentMonth == 1 {
		monthName = "Enero"
	} else if currentMonth == 2 {
		monthName = "Febrero"
	} else if currentMonth == 3 {
		monthName = "Marzo"
	} else if currentMonth == 4 {
		monthName = "Abril"
	} else if currentMonth == 5 {
		monthName = "Mayo"
	} else if currentMonth == 6 {
		monthName = "Junio"
	} else if currentMonth == 7 {
		monthName = "Julio"
	} else if currentMonth == 8 {
		monthName = "Agosto"
	} else if currentMonth == 9 {
		monthName = "Septiembre"
	} else if currentMonth == 10 {
		monthName = "Octubre"
	} else if currentMonth == 11 {
		monthName = "Noviembre"
	} else if currentMonth == 12 {
		monthName = "Diciembre"
	}
	return monthName
}

func monthNameForNumberV2(currentMonth int) string {

	var monthName string

	switch currentMonth {
		case 1:	 monthName = "Enero"
		case 2:	 monthName = "Febrero"
		case 3:	 monthName = "Marzo"
		case 4:	 monthName = "Abril"
		case 5:	 monthName = "Mayo"
		case 6:	 monthName = "Junio"
		case 7:	 monthName = "Julio"
		case 8:	 monthName = "Agosto"
		case 9:	 monthName = "Septiembre"
		case 10: monthName = "Octubre"
		case 11: monthName = "Noviembre"
		case 12: monthName = "Diciembre"
	}
	return monthName
}

func monthNameForNumberV3(currentMonth int) string {

	var months = [12]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}	

	return months[currentMonth - 1]
}

func monthNameForNumberV4(currentMonth int) string {

	var months = map[int]string{
		1:	"Enero",
		2:	"Febrero",
		3:	"Marzo",
		4:	"Abril",
		5:	"Mayo",
		6:	"Junio",
		7:	"Julio",
		8:	"Agosto",
		9:	"Septiembre",
		10:	"Octubre",
		11:	"Noviembre",
		12:	"Diciembre",
	}

	return months[currentMonth]
}
