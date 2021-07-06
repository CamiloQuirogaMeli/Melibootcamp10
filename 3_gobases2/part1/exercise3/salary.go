package main

import (
	e "errors"
	f "fmt"
	"github.com/leekchan/accounting"
)

const (
	SALARY_HOUR_A float32 = 3000
	SALARY_HOUR_B float32 = 1500
	SALARY_HOUR_C float32 = 1000
	INCENTIVE_A   float32 = 1.20
	INCENTIVE_B   float32 = 1.50
)

func salary(workMinutes float32, category string) (string, error) {
	var workHours float32
	var salary float32

	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	workHours = workMinutes / 60

	switch {
	case category == "A" || category == "a":
		salary = SALARY_HOUR_A * workHours * INCENTIVE_A
	case category == "B" || category == "b":
		salary = SALARY_HOUR_B * workHours * INCENTIVE_B
	case category == "C" || category == "c":
		salary = SALARY_HOUR_C * workHours ¿1000
	default:
		return ac.FormatMoney(0), e.New("please, enter a valid category (A, B or C)")
	}

	return ac.FormatMoney(salary), nil
}

func main() {
	var workMinutes float32
	var category string

	f.Print("Welcome, enter your category: ")
	f.Scanln(&category)
	f.Print("Welcome, enter your minutes worked in last month: ")
	f.Scanln(&workMinutes)

	res, err := salary(workMinutes, category)

	if err == nil {
		f.Println("Your salary is:", res)
	} else {
		f.Println(err)
	}
}
