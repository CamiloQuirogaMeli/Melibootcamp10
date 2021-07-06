package tm

func TaxPercentage(salary float64) (tax float64) {
	if salary > 50000 {
		tax = 0.17
	} else if salary > 150000 {
		tax = 0.10
	} else {
		tax = 0
	}
	return
}
