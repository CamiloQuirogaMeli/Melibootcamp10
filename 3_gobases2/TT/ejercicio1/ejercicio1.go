package main

import "fmt"

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func (s Student) detail() {
	fmt.Printf("Alumno %s %s con identificacion %d\nIngreso fue %s", s.Name, s.LastName, s.DNI, s.Date)
}

func main() {
	a := Student{"Carlos", "Infante", 10108274635, "20/1/2017"}
	a.detail()
}
