package main

import "fmt"

func main() {
	// Arrays
	var names [2]string
	names[0] = "Giorno"
	names[1] = "Giovanna"

	fmt.Printf("Printing first position of array: %s\n", names[0])

	// Slice declaration
	var cats = []string{"Daphne", "Fred"}

	// To access value of an slice, just the index left
	fmt.Printf("Printing the first cat in my slice: %s\n", cats[0])

	// Crear slice con make()
	dogs := make([]string, 3)
	println(len(dogs))
	dogs[0] = "Luna"
	dogs[1] = "Lupe"
	fmt.Println("Printing the content of the slice: ", dogs)

	// Range (Like Python)
	fmt.Println(dogs[1:3], "Printing slice range. Arrays also can do this")

	// Slices can receive more data through append()
	dogs = append(dogs, "Renato")

	fmt.Println("Appended ", dogs)

	// Maps
	/* First assignment works as type of keys. Second is type of values  */
	var family = map[string]string{}

	//  Make function can declare maps, too
	friends := make(map[string]int)

	// Len function returns 0 to an uninitialized map
	fmt.Println(len(friends), "<- Should print 0")
	fmt.Println(len(family), "<- Should print 0")

	// Adding elements to a non initialized map. Update works like that
	friends["Alejandro"] = 23
	fmt.Println(friends)

	// Delete elements with delete method
	friends["Juan"] = 22
	fmt.Println(friends, "With Juan")
	delete(friends, "Juan")
	fmt.Println(friends, "Without Juan")

	// Iterate through maps
	for key, value := range friends {
		fmt.Println("Key:", key, "| Value:", value)
	}

	// If and short declarations
	if age := 33; age%3 == 0 {
		fmt.Println("3x multiple")
	} else {
		fmt.Println("Not 3x multiple")
	}

	// Switch without condition
	age := 33
	switch {
	case age < 18:
		fmt.Println("Not adult")
	case age >= 18:
		fmt.Println("Adult")
	case age > 100:
		fmt.Println("Too old")
	}

	// Also, switch can be declared like this:
	switch day := "Sunday"; day {
	case "Lunes", "Martes", "Miércoles", "Jueves", "Viernes":
		fmt.Println("Weekday")
	default:
		fmt.Println("It's weekend")
	}

	// Switch have fallthrough option to execute the next set of cases for the same sentence
	switch day := 27; {
	case day%7 == 0:
		fmt.Println("Multiple of 7")
	case day%3 == 0:
		fmt.Println("Multiple of three")
		fallthrough
	case day%7 == 0 && day%3 == 0:
		fmt.Println("Multiple of both!")
	default:
		fmt.Println("Normal")
	}

	// While :y
	i := 23
	for i > 20 {
		fmt.Println("I'm fine")
		i--
	}

}
