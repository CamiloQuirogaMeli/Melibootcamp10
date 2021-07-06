package main

func salaryTaxes(net int) (tax float64) {
	switch {
	case net > 50000:
		tax = float64(net) * 0.17
		fallthrough
	case net > 150000:
		additional := float64(net) * 0.10
		tax += additional
	default:
		tax = 0
	}

	return tax
}
