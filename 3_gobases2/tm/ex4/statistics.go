package main

import (
  "fmt"
  "errors"
)

func statistics(scores ...int) (min, max int, avg float64, err error){
  var sum, minim, maxim int
  var ave float64
  minim = scores[0]
  maxim = scores[0]
  for i:=0; i<len(scores); i++{
    if scores == nil{
      return 0, 0, 0, errors.New("No es posible el calculo")
    }
    if scores[i]<minim{
      minim = scores[i]
    }
    if scores[i]>maxim{
      maxim = scores[i]
    }
  }
  for _, value := range scores{
    sum += value
  }
  ave = float64(sum)/float64(len(scores))

  return minim, maxim, ave, nil
}

func main() {
  min, max, avg, err := statistics(10,2,33,4,5)

  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println(min, max, avg)
  }
}
