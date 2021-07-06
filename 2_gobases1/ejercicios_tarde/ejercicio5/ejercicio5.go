package main
import "fmt"

func main() {
	var nuevoEstudiante string
	var estudiantes = []string{"Benjamin", 
				"Nahuel", 
				"Brenda", 
				"Marcos", 
				"Pedro", 
				"Axel", 
				"Alez", 
				"Dolores", 
				"Federico", 
				"Hernan", 
				"Leandro", 
				"Eduardo", 
				"Duvraschka"}

	fmt.Println("Lista de estudiantes:", estudiantes)
	fmt.Println("Ingrese el nombre del nuevo estudiante: ")
	fmt.Scanln(&nuevoEstudiante)

	estudiantes = append(estudiantes, nuevoEstudiante)
	fmt.Println("Lista actualizada de estudiantes:", estudiantes)
}
