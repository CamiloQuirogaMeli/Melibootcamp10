package main

import "fmt"

func main() {
// a) --------------------------
	var students = []string {
		"Benjamin","Nahuel","Brenda","Marcos","Pedro","Axel","Alez","Dolores","Federico","Hernan","Leandro","Eduardo","Duvraschka",
	}
	fmt.Println("Estudiantes de la primera clase:")
	printStudents(students)
// b) --------------------------
	students = append(students, "Gabriela")
	fmt.Println("Estudiantes luego de dos clases:")
	printStudents(students)
}

func printStudents(students []string) {
	for i := 0; i < len(students); i++ {
		fmt.Printf("%d.\t%s\n", i+1, students[i])
	}
}
