package main

func calcularImpuesto(sueldo float64) float64 {

	if sueldo > 15000 {
		return (sueldo * 0.27)
	} else if sueldo >= 50000 {
		return (sueldo * 0.17)
	}
	return 0
}
