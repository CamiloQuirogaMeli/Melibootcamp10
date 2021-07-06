package main

import "fmt"

//elegi hacerlo con map por su facilidad para buscar datos a traves de la clave. Se puede hacer tambien con switch
func main() {
	var meses = map[int]string{}

	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"

	fmt.Print("Ingrese el mes: ")
	var mes int
	fmt.Scan(&mes)

	fmt.Println(mes, ",", meses[mes])

}
