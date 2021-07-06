package main

import "fmt"

func main() {
	var employes = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Benjamin es: ", employes["Benjamin"])
	fmt.Println("Los empleados mayores a 21 son:")
	for key, element := range employes {
		if element > 21 {
			fmt.Println(key)
		}
	}
	fmt.Println("Incorporando a Federico a la empresa")
	employes["Federico"] = 25
	fmt.Println("Borrando a Pedro de la lista de empleados, lo siento Pedro :(")
	delete(employes, "Pedro")
	fmt.Println("Nueva lista de empleados:")
	for element := range employes {
		fmt.Print(element, " ")
	}
}
