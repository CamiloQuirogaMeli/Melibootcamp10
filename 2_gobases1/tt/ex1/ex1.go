package main

import ("fmt")

func main(){
  var word string

  fmt.Printf("Palabra: ")
  fmt.Scanf("%s", &word)

  fmt.Printf("La palabra %s tiene una cantidad de %d letras\n", word, len(word))

  for i, letter := range word {
    fmt.Printf("Letra %d: %c\n", i+1, letter)
  }
}
