package main

import (
	"errors"
	"fmt"
)

type Empleado struct {
	horas    int
	salary   float64
	salaries []float64
}

func (e *Empleado) calcSalary() (int, float64, error) {
	if e.horas < 80 || e.horas < 0 {
		return e.horas, e.salary, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if e.salary >= 150000 {
		e.salary *= 0.9
	}
	return e.horas, e.salary, nil
}

func (e *Empleado) calcAguinaldo() (float64, error) {
	var best float64
	if e.salary < 0 {
		return 0, errors.New("error: ingresó un numero negativo")
	}
	for _, month := range e.salaries {
		if best < month {
			best = month
		}
	}
	agui := (best / 12) * float64(len(e.salaries))
	return agui, nil
}

func main() {
	fmt.Println("vim-go")
	e := Empleado{horas: 90, salary: 150000, salaries: []float64{39000, 70000, 50000, 15900, 43000, 32000}}

	h, s, err := e.calcSalary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tu sueldo despues de impuestos es de %.2f con un total de horas trabajadas de: %d\n", s, h)
	}
	a, err2 := e.calcAguinaldo()
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tu aguinaldo por medio año es de: %.2f\n", a)
	}

}
