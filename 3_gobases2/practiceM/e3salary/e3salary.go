package main

import "fmt"

func main() {

	res := salary("C", 3600)
	fmt.Println(res)

}

func salary(category string, minutes int) float64 {
	var aux float64
	switch category {
	case "A":
		{
			aux = (3000.0 * float64(minutes)) * 1.50 / 3600.0
		}
	case "B":
		{
			aux = (1500 * float64(minutes)) * 1.2 / 3600
		}
	case "C":
		{
			aux = (1000 * float64(minutes)) / 3600
		}
	}
	return aux
}
