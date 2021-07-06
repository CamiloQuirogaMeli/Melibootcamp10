package ejercicio1

import (
	"strconv"
	"testing"
)

func TestStudentDetails(t *testing.T) {
	student := Student{
		Name:    "Marlon",
		LasName: "Barreto",
		DNI:     123,
		Date: date{
			day:   15,
			month: 06,
			year:  2021,
		},
	}
	expectedString := "Nombre: [" + student.Name + "]\n" +
		"Apellido: [" + student.LasName + "]\n" +
		"DNI: [" + strconv.Itoa(student.DNI) + "]\n" +
		"Fecha: [" + strconv.Itoa(student.Date.day) +
		"-" + strconv.Itoa(student.Date.month) +
		"-" + strconv.Itoa(student.Date.year) + "]\n"
	detailsString := student.details()
	if detailsString != expectedString {
		t.Error("Test Failed: {} expected, {} received\n", expectedString, "\n", detailsString)
	}
}
