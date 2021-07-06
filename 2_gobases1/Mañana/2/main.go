package main

import "fmt"

func main() {
	var temperature string
	var humidity string
	var pressure string

	temperature = "12.6 ºC"
	humidity = "27%"
	pressure = "975.3 hPa"

	fmt.Println("Temperatura :", temperature)
	fmt.Println("Humedad :", humidity)
	fmt.Println("Presion :", pressure)
}

/* Respuesta punto 3:
En caso que se quieran mostrar las unidades utilizo string, en caso que solo se muestren los valores usaria el tipo float.
*/
