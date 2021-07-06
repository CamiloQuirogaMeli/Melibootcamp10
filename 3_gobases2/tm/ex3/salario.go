package main

import (
  "fmt"
)

func calcSalary(min int, category string) float64 {
  hours := min/60
  var salary float64
  switch category {
    case "C":
      salary = float64(hours) * 1000
    case "B":
      salary = float64(hours) * 1500
      salary *= 1.2
    case "A":
      salary = float64(hours) * 3000
      salary *= 1.5
  }
  return salary
}

func main() {
  fmt.Println(calcSalary(9600, "B"))
}
