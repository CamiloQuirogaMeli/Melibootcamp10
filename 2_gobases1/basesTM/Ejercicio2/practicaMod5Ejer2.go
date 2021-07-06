package main

import "fmt"

func main() {
	var temp int
	var humedad int
	var presion float64
	temp = 17
	humedad = 48
	presion = 1016.6
	fmt.Println("temperatura: ", temp, "º\nhumedad: ", humedad, "%\npresion: ", presion, "mb")

}

//tipo datos como puse int para temperatura y humedad y float para presion ya que puede contener numeros fraccionales
