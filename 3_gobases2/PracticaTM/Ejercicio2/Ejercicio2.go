// Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
// devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

package main

import (
	"errors"
	"fmt"
	"strings"
)

func calculateProm(studentQualifications []float64) (float64, error) {

	var prom float64 = 0

	for _, qualification := range studentQualifications {
		if qualification <= 0 {
			return 0, errors.New("no se pueden ingresar números negativos")
		}
		prom += qualification
	}
	prom /= float64(len(studentQualifications))
	return prom, nil
}

func main() {

	var studentQualifications []float64
	var qualification float64
	var insertQualificationStudent string = "si"

	for strings.Compare(insertQualificationStudent, "si") == 0 {
		fmt.Printf("Ingrese la nota del alumno: ")
		fmt.Scan(&qualification)
		studentQualifications = append(studentQualifications, qualification)
		fmt.Printf("¿Desea ingresar otra nota?(si/no): ")
		fmt.Scan(&insertQualificationStudent)
		insertQualificationStudent = strings.ToLower(insertQualificationStudent)
	}

	res, err := calculateProm(studentQualifications)

	if err == nil {
		fmt.Printf("El promedio del estudiante es: %.2f\n", res)
	} else {
		fmt.Printf("Error en el ingreso de notas\n")
	}

}
