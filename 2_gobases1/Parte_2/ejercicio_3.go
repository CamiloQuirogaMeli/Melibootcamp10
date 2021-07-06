package main

import "fmt";

const EDAD_MINIMA = 22
const SUELDO_MINIMO = 100000
const ANTIGUEDAD_MINIMA = 1

var (
  edad uint8= 75
  tieneEmpleo = true
  sueldo = 999999
  aniosAntiguedad = 0
)

func main() {
  if edad < EDAD_MINIMA {
    fmt.Println("No tiene la edad mínima")
  } else {
    fmt.Println("Cumple con la edad!")
  }
  if sueldo > SUELDO_MINIMO {
    fmt.Println("No tiene el sueldo suficiente")
  } else {
    fmt.Println("Cuenta con el sueldo necesario")
  }
  if tieneEmpleo {
    if aniosAntiguedad >= ANTIGUEDAD_MINIMA {
      fmt.Println("Tiene la antiguedad suficiente")
    } else {
      fmt.Println("No cuenta con los años de antiguedad mínima")
    }
  } else {
    fmt.Println("No tiene empleo")
  }
}
