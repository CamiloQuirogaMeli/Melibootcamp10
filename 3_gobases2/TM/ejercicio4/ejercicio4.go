package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	MIN = "min"
	MAX = "max"
	AVG = "avg"
)

func minVal(grades ...int) float64 {
	var minGrade = math.MaxInt64
	for _, val := range grades {
		if val < minGrade {
			minGrade = val
		}
	}
	return float64(minGrade)
}

func maxVal(grades ...int) float64 {
	var maxGrade = math.MinInt64
	for _, val := range grades {
		if val > maxGrade {
			maxGrade = val
		}
	}
	return float64(maxGrade)
}

func average(grades ...int) float64 {
	valuesSum, students := 0, len(grades)
	for _, val := range grades {
		valuesSum += val
	}
	return float64(valuesSum / students)
}

func studentStatistics(stat string) (func(...int) float64, error) {
	switch stat {
	case MIN:
		return minVal, nil
	case MAX:
		return maxVal, nil
	case AVG:
		return average, nil
	default:
		return nil, errors.New("No se pudo calcular ninguna estadistica")
	}
}

func main() {
	minFunc, err := studentStatistics(MIN)
	maxFunc, err := studentStatistics(MAX)
	avgFunc, err := studentStatistics(AVG)
	if err != nil {
		fmt.Println(err)
	} else {
		minValue := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		avgValue := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Valor minimo es %.2f\n", minValue)
		fmt.Printf("Valor maximo es %.2f\n", maxValue)
		fmt.Printf("Valor promedio es %.2f\n", avgValue)
	}

}
