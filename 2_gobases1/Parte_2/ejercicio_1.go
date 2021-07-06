package main

import "fmt";

var (
  palabra = "Luis Alberto Oropeza Vilchis"
)

func main() {
  fmt.Println("La longitud de la palabra es:", len(palabra))
  for i:=0; i < len(palabra); i++ {
    fmt.Println(string(palabra[i]))
  }
}
