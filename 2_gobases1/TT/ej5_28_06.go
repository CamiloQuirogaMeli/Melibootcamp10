package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s []string
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println(students)

	fmt.Println("Ingrese otro estudiante")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	s = append(students, text)
	fmt.Printf("%s", s)

}
