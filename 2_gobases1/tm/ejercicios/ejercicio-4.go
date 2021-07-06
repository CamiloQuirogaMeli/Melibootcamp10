package ejercicios

import "fmt"

func CuartoEjercicio(mes int) {
	var texto string
	if mes == 1 {
		texto = "Enero"
	} else if mes == 2 {
		texto = "Febrero"
	} else if mes == 3 {
		texto = "Marzo"
	} else if mes == 4 {
		texto = "Abril"
	} else if mes == 5 {
		texto = "Mayo"
	} else if mes == 6 {
		texto = "Junio"
	} else if mes == 7 {
		texto = "Julio"
	} else if mes == 8 {
		texto = "Agosto"
	} else if mes == 9 {
		texto = "Septiembre"
	} else if mes == 10 {
		texto = "Octubre"
	} else if mes == 11 {
		texto = "Noviembre"
	} else if mes == 12 {
		texto = "Diciembre"
	} else {
		texto = "Número de mes no válido"
	}

	fmt.Println(texto)
}

func CuartoEjercicioVersionDos(mes int) {
	var texto string
	switch mes {
	case 1:
		texto = "Enero"
	case 2:
		texto = "Febrero"
	case 3:
		texto = "Marzo"
	case 4:
		texto = "Abril"
	case 5:
		texto = "Mayo"
	case 6:
		texto = "Junio"
	case 7:
		texto = "Julio"
	case 8:
		texto = "Agosto"
	case 9:
		texto = "Septiembre"
	case 10:
		texto = "Octubre"
	case 11:
		texto = "Noviembre"
	case 12:
		texto = "Diciembre"
	default:
		texto = "Número de mes no válido"
	}

	fmt.Println(texto)
}

func CuartoEjercicioVersionTres(mes int) {
	var meses = [12]string{
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
		"Diciembre"}
	if 1 <= mes && mes <= 12 {
		fmt.Println(meses[mes-1])
	} else {
		fmt.Println("Número de mes no válido")
	}
}
