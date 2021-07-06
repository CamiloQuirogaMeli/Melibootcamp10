package main

import "fmt"

func main()  {
	sueldol := salario(500000)
	fmt.Println(sueldol)
}
func salario (sueldo float64)  float64 {

	switch {
	case sueldo < 50000:
		return sueldo
	case sueldo > 50000:
		sueldoFinal := sueldo - (sueldo * 0.17)
		return sueldoFinal
	case sueldo > 150000:
		sueldoFinal := sueldo - (sueldo * 0.27)
		return sueldoFinal
	}

	return sueldo
}