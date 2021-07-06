package main

import "fmt"

func main() {
	var temperatura float32 = 3.2
	var humedad float32 = 88
	var presion float32 = 1010

	//fmt.Printf("La temperatura es de %.2f ºC, la humedad %.2f %, y la presion %.2f Hp \n", temperatura, humedad, presion)
	fmt.Println("La temperatura es de", temperatura, "ºC, la humedad ", humedad, "%, y la presion ", presion, "Hp")
}
