package main

import "fmt"

const (
  SALARIO_IMPUESTO_BASE = 50000.0
  SALARIO_IMPUESTO_MAYOR = 150000.0
  IMPIESTO_BASE = 0.17
  IMPUESTO_EXTRA = 0.10
)

func main() {
  salario := 200000.0
  impuesto := calcularImpuesto(salario)
  fmt.Println("Impuestos:", impuesto)
}

func calcularImpuesto(salario float64) float64 {
  var impuesto float64

  if salario > SALARIO_IMPUESTO_MAYOR {
    impuesto = salario * (IMPIESTO_BASE + IMPUESTO_EXTRA)
  } else if salario > SALARIO_IMPUESTO_BASE {
    impuesto = salario * IMPIESTO_BASE
  } else {
    impuesto = 0
  }
  return impuesto
}
