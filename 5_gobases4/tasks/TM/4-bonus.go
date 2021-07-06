package main

import "errors"

func bonus(semesterSalary []int) (bonus int, err error) {
	max, err := maxSalary(semesterSalary)

	if err != nil {
		return 0, err
	}

	bonus = (max * 12) / len(semesterSalary)

	return bonus, nil
}

func maxSalary(semesterSalary []int) (max int, err error) {
	if len(semesterSalary) == 0 {
		return 0, errors.New("error: semester salaries should be provided")
	} else if len(semesterSalary) > 6 {
		return 0, errors.New("error: la cantidad de meses trabajados es mayor a 6")
	}

	for _, month := range semesterSalary {
		if month < 0 {
			return 0, errors.New("error: el valor de un salario no puede ser negativo")
		}
		if month > max {
			max = month
		}
	}

	return max, nil
}
