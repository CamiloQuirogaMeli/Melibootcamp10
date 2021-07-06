package main

import "fmt"

func main()  {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	cant := 0

	for _ , e := range employees{
      if e > 21{
      	cant += 1
	  }
	}

	fmt.Println("Existen", cant, "empleados mayores de 21 años" )

}