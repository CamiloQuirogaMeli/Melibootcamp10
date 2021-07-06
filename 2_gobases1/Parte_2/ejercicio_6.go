package main

import "fmt";

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
  contador := 0
  fmt.Println("Mapa inicial:", employees)
  fmt.Println("La edad de Benjamin es:", employees["Benjamin"])
  for _, valor := range employees {
    if valor > 21 {
      contador++
    } 
  }
  fmt.Println("Empleados mayores a 21:", contador)
  employees["Federico"] = 25
  fmt.Println("Después de agregar:", employees)
  delete(employees, "Pedro")
  fmt.Println("Después de eliminar:", employees)
  
}
