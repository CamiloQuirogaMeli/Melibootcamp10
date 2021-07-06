package main

import "fmt"

func main() {
  estudiantes := []string{
    "Luis",
    "Nahuel",
    "Brenda",
    "Marcos",
    "Pedro",
    "Axel",
    "Alez",
    "Dolores",
    "Federico",
    "Hernan",
    "Leandro",
    "Eduardo",
    "Duvraschka",
  }
  fmt.Println(estudiantes)
  estudiantes = append(estudiantes, "Gabriela")
  fmt.Println(estudiantes)
}
