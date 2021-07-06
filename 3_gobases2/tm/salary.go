package tm

import (
	"strings"
)

const (
	A = "A"
	B = "B"
	C = "C"
)

func CalculateSalary(minutes uint, category string) (salary float64) {
	hours := minutesToHours(minutes)
	computeSalary := salaryType(category)
	salary = computeSalary(hours)
	return
}

func minutesToHours(minutes uint) (hours float64) {
	hours = float64(minutes)/60
	return
}

func salaryType(salaryType string) func(hours float64) (salary float64) {
	strings.ToUpper(salaryType)
	switch salaryType {
	case A:
		return salaryA
	case B:
		return salaryB
	case C:
		return salaryC
	default:
		return nil
	}
}

func salaryA(hours float64) (salary float64) {
	salary = 3000 * hours
	salary += salary*0.5
	return
}

func salaryB(hours float64) (salary float64) {
	salary = 1500 * hours
	salary += salary*0.2
	return
}

func salaryC(hours float64) (salary float64) {
	salary = 1000 * hours
	return
}