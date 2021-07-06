package main

import "fmt";

var (
  precio = 75
  descuento = 0.12
)

func main() {
  nuevoPrecio := float64(precio) * (1 - descuento)
  fmt.Println("El precio con descuento es:", nuevoPrecio)
}
