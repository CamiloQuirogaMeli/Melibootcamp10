package main

import (
  "fmt"
  "errors"
)

func calcAverage(scores ...float64) (float64, error) {
  var suma, result float64

  for _, value := range scores{
    if value<0{
      return 0, errors.New("Las calificaciones deben ser positivas")
    }
    suma += value
  }
  result = suma/float64(len(scores))

  return result, nil
}

func main() {
  res, err := calcAverage(2,-5,8,10)

  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println(res)
  }
}
