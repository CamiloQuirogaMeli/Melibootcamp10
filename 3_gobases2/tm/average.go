package tm

import "errors"

func Average(grades ...int) (average float64, err error) {
	for _, grade := range grades {
		if grade >= 0 {
			average += float64(grade)
		} else {
			err = errors.New("Negative grades are not accepted")
		}
	}
	average /= float64(len(grades))
	return

}
