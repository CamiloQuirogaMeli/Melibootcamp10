package main

import "fmt";

var (
  temperatura int8 = 17
  humedad uint8 = 73
  presion int = 1019
)

func main() {
  fmt.Printf("Temperatura: %v Tipo: %T\n", temperatura, temperatura)
  fmt.Printf("Humedad: %v Tipo: %T\n", humedad, humedad)
  fmt.Printf("Presion: %v Tipo: %T\n", presion, presion)
}
