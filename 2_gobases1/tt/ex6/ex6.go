package main

import (
  "fmt"
)

func main(){
  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

  fmt.Println(employees)

  fmt.Printf("Benjamin tiene %d años.\n", employees["Benjamin"])

  employees["Francisco"] = 25

  delete(employees, "Pedro")

  fmt.Println(employees)
}
