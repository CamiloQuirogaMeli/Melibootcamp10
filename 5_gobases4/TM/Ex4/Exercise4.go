package main

import (
	"errors"
	"fmt"
)

func calcSalary(hours int) (float64, error) {
	if hours < 80 {
		return 0, errors.New("the worker cannot have worked less than 80 hours per month")
	}

	salary := float64(hours * 1500)

	if salary > 150000 {
		salary *= 0.9
	}

	return salary, nil
}

func calcBonus(months int, bestSalary float64) (float64, error) {

	if months < 0 || bestSalary < 0 {
		return 0, errors.New("months or salary can't be negatives")
	}

	bonus := float64(bestSalary/12) * float64(months)

	return bonus, nil

}

func main() {

	salary, salErr := calcSalary(80)

	if salErr != nil {
		err := fmt.Errorf("error calculating salary: %w", salErr)

		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("Salary:", salary)
	}

	bonus, bonusErr := calcBonus(6, 12000.00)

	if bonusErr != nil {
		err := fmt.Errorf("error calculating  bonus: %w", bonusErr)

		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("Bonus:", bonus)
	}

}
