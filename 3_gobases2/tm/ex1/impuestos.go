package main

import (
  "fmt"
)

func impuestos(sueldo float64) float64{
  if sueldo > 50000 && sueldo < 150000{
    return sueldo*0.17
  }
  if sueldo > 150000{
    return sueldo*0.27
  }
  return 0
}

func main() {
  var sueldo, impuesto float64
  fmt.Scanln(&sueldo)
  impuesto = impuestos(sueldo)
  fmt.Printf("Los impuestos a generar son de: %.2f\n", impuesto)
}
