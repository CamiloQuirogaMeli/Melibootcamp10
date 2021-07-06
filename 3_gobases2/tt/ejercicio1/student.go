package ejercicio1

import (
	"strconv"
	"strings"
)

type Student struct {
	Name    string
	LasName string
	DNI     int
	Date    date
}

type date struct {
	day   int
	month int
	year  int
}

func (s Student) details() string {
	var sb strings.Builder
	sb.WriteString("Nombre: [")
	sb.WriteString(s.Name)
	sb.WriteString("]\n")
	sb.WriteString("Apellido: [")
	sb.WriteString(s.LasName)
	sb.WriteString("]\n")
	sb.WriteString("DNI: [")
	sb.WriteString(strconv.Itoa(s.DNI))
	sb.WriteString("]\n")
	sb.WriteString("Fecha: [")
	sb.WriteString(strconv.Itoa(s.Date.day))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(s.Date.month))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(s.Date.year))
	sb.WriteString("]\n")
	return sb.String()
}
