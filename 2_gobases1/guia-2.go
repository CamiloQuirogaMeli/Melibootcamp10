package main

import "fmt"

func main() {
	//var word string = "bootcamp"

	//var price float64 = 500.5
	//var discountPercent float64 = 20.0

	//lettersOfAWord(word)
	//discount(price, discountPercent)
	//caseCredits()
	//months(3)
	namesList()
	age()
}

func lettersOfAWord(word string) {

	fmt.Println("The word has ", len(word), " letters.")
	fmt.Println("")
	y := 1
	for i := 0; i < len(word); i++ {
		fmt.Println(word[i:y])
		y++
	}
}

func discount(price float64, discount float64) {

	priceWithDiscount := price - ((discount / 100) * price)

	fmt.Println("The price with discount is: ", priceWithDiscount)
}

func caseCredits() {
	var employed bool
	var age int
	var seniorityYears int
	var salary float64

	employed = true
	age = 24
	seniorityYears = 2
	salary = 120000

	switch {
	case !employed:
		fmt.Println("El cliente actualmente no está empleado.")
	case age < 22:
		fmt.Println("El cliente es menor de 22 años.")
	case salary < 100000:
		fmt.Println("Al cliente se le cobrarán intereses.")
	case seniorityYears < 1:
		fmt.Println("El cliente debe tener al menos 1 año de antigüedad.")
	default:
		fmt.Println("El préstamo pueder ser otorgado correctamente.")
	}
}

func months(month int) {
	switch month {
	case 1:
		fmt.Println(month, ".January")
	case 2:
		fmt.Println(month, ".February")
	case 3:
		fmt.Println(month, ".March")
	case 4:
		fmt.Println(month, ".April")
	case 5:
		fmt.Println(month, ".May")
	case 6:
		fmt.Println(month, ".June")
	case 7:
		fmt.Println(month, ".Julie")
	case 8:
		fmt.Println(month, ".August")
	case 9:
		fmt.Println(month, ".September")
	case 10:
		fmt.Println(month, ".October")
	case 11:
		fmt.Println(month, ".November")
	case 12:
		fmt.Println(month, ".December")
	default:
		fmt.Println(month, ". character invalid")
	}
	/* este ejercicio se podría haber hecho de otras maneras pero la forma que elegí me perece
	la más sencilla y más legible
	Se podría haber hecho con una sucesión de IF...ESLE IF,
	con un Map previamente cargado y luego obtener el mes con una clave numérica
	*/
}

func namesList() {

	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(students)
	//item 2
	students = append(students, "Gabriela")
	fmt.Println(students)
}

func age() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var olderEmployees []string
	var count int = 0
	//item 1
	fmt.Println("La edad de Benjamin es ", employees["Benjamin"])
	//item 2
	for key, value := range employees {
		if value > 21 {
			count++
			olderEmployees = append(olderEmployees, key)
		}
	}
	fmt.Println("Los empleados mayores a 21 años son: ", count, olderEmployees)
	//item 3
	employees["Federico"] = 25
	fmt.Println(employees)
	//item 4
	delete(employees, "Pedro")
	fmt.Println(employees)
}
