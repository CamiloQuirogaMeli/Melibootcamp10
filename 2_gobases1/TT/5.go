package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* a. Una profesor de la universidad quiere tener un listado con todos sus estudiantes. Es
necesario generar una aplicación que contenga dicha lista.
Estudiantes:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernan,
Leandro, Eduardo, Duvraschka.

b. Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
sin modificar el código que escribiste inicialmente.
Estudiante:
Gabriela */

var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan",
"Leandro", "Eduardo", "Duvraschka"}

func main() {
	listadoDeNombres()
}

func listadoDeNombres() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Agrega un estudiante a la lista")
	fmt.Println("Input format: name, ex:")
	fmt.Println("-> Gonzalo")
	fmt.Println("---------------------")

	for {
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  // convert CRLF to LF (Only Windows)
	  text = strings.Replace(text, "\n", "", -1)
	  
	  if len(text) > 0 { students = append(students, text) }
	  fmt.Println("Lista de estudiantes:", students)
	}
}