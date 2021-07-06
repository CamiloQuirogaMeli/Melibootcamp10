package main

import "fmt"

type BeneficiosClase struct {
  salarioPorMinuto float64
  extra float64
}

var clases =  map[string] BeneficiosClase {
  "C": {1000.0/60.0, 0.0},
  "B": {1500.0/60.0, 0.20},
  "A": {3000.0/60.0, 0.50},
}


func main() {
  salario := calcularSalario("B", 60)
  fmt.Printf("El salario es: %.2f\n", salario)
}

func calcularSalario(categoria string, minutosTrabajados int) float64 {
  beneficios := clases[categoria]
  salario := float64(minutosTrabajados) * beneficios.salarioPorMinuto
  salario += salario * beneficios.extra
  return salario
}
