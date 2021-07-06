package main

import ("fmt")

func main(){
  var (
    edad int
    empleado string
    antiguedad int
    sueldo float64
  )

  fmt.Print("Edad: ")
  fmt.Scanf("%d\n", &edad)
  fmt.Print("Empleado: ")
  fmt.Scanf("%s\n", &empleado)
  fmt.Print("Meses de antigüedad: ")
  fmt.Scanf("%d\n", &antiguedad)
  fmt.Print("Sueldo: ")
  fmt.Scanf("%f\n", &sueldo)

  if edad<22 || empleado=="no" || antiguedad<12{
    fmt.Println("\nNo es posible el préstamo.")
  }else if sueldo>=100000{
    fmt.Println("\nPréstamo otorgado. No se te cobrarán interéses.")
  }else{
    fmt.Println("\nPréstamo otorgado. Se te cobrarán interéses.")
  }
}
