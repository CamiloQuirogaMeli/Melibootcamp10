package main

import "fmt"

func main() {
  promedio := calcularPromedio(4, 5, 4, 5, 6, -1) 
  fmt.Println("El promedio es:", promedio)
}

func calcularPromedio(calificaciones ...int) float64 {
  var total float64
  for _, valor := range calificaciones {
    if valor < 0 {
      fmt.Println("Error, no puede tener calificaciones negativas.")
      return 0
    }
    total += float64(valor)
  }
  return total / float64(len(calificaciones))
}
