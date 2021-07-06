package main

import (
  "fmt"
)

func main(){
  var (
    temperature float64 = 19.5
    humidity string = "57%"
    preasure float64 = 1027.4
  )

  fmt.Print("Temperatura: ")
  fmt.Println(temperature)
  fmt.Print("Humedad: ")
  fmt.Println(humidity)
  fmt.Print("Presion: ")
  fmt.Println(preasure)
}
