package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	min = "min"
	avg = "avg"
	max = "max"
)

func convertTo64(ar []int) []float64 {
	nums := make([]float64, len(ar))
	for i, v := range ar {
		nums[i] = float64(v)
	}
	return nums
}

func calcMin(nums ...int) (float64, error) {
	nums64 := convertTo64(nums)
	var min float64 = nums64[0]
	for _, n := range nums64 {
		if n < 0 {
			return 0, errors.New("no negative numbers allowed")
		}
		if n < min {
			min = n
		}
	}
	return min, nil
}

func calcMax(nums ...int) (float64, error) {
	nums64 := convertTo64(nums)
	var max float64 = nums64[0]
	for _, n := range nums64 {
		if n < 0 {
			return 0, errors.New("no negative numbers allowed")
		}
		if n > max {
			max = n
		}
	}
	return max, nil
}

func calcAvg(nums ...int) (float64, error) {
	nums64 := convertTo64(nums)
	var total float64 = 0
	for _, n := range nums64 {
		if n < 0 {
			return 0, errors.New("no negative numbers allowed")
		}
		total += n
	}
	avg := total / float64(len(nums64))
	return avg, nil
}

func operation(function string) func(nums ...int) (float64, error) {
	switch function {
	case min:
		return calcMin
	case avg:
		return calcAvg
	case max:
		return calcMax
	}
	return nil
}

func main() {
	op := operation(max)

	valorMinimo, err := op(2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(valorMinimo)
}
