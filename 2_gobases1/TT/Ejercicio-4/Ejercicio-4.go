package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Ejercicio 4")
	rand.Seed(time.Now().UnixNano())
	var max, min int = 12, 1
	var month int = rand.Intn(max-min) + min // Obtener un numero random en el rango 1-12
	mapMonth := make(map[int]string)
	mapMonth[1] = "Enero"
	mapMonth[2] = "Febrero"
	mapMonth[3] = "Marzo"
	mapMonth[4] = "Abril"
	mapMonth[5] = "Mayo"
	mapMonth[6] = "Junio"
	mapMonth[7] = "Julio"
	mapMonth[8] = "Agosto"
	mapMonth[9] = "Septiembre"
	mapMonth[10] = "Octubre"
	mapMonth[11] = "Noviembre"
	mapMonth[12] = "Diciembre"
	fmt.Println("El mes correspondiente al numero", month, "es: ", mapMonth[month])

}
