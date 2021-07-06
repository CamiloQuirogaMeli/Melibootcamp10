package tm

import (
	"errors"
	"sort"
)

const (
	MIN = "minimo"
	AVG = "promedio"
	MAX = "maximo"
)

func Statistics(operation string) func (grades ...int) (int, error) {
	switch operation {
	case MIN:
		return min
	case MAX:
		return max
	case AVG:
		return avg
	default:
		return nil
	}
}

func min(grades ...int) (int, error) {
	sort.Ints(grades)
   	if grades[0] < 0 {
   		return grades[0],errors.New("invalid argument")
	}
	return grades[0], nil
}

func max(grades ...int) (int, error) {
	sort.Ints(grades)
	lastPosition := len(grades)-1
	if grades[lastPosition] < 0 {
		return 0,errors.New("invalid argument")
	}
	return grades[lastPosition],nil
}

func avg(grades ...int) (int, error) {
	average,err := Average(grades...)
	return int(average),err
}