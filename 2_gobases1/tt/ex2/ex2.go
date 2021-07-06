package main

import ("fmt")

func main(){
  var precio, descuento float64

  descuento = 0.75

  fmt.Printf("Ingresa el precio de tu producto: ")
  fmt.Scanf("%f", &precio)

  precio *= descuento

  fmt.Printf("Tu producto tiene el %.2f%% de descuento\n", descuento*100)
  fmt.Printf("Precio final: $%.2f\n", precio)
}
