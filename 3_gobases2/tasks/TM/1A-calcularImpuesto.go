package main

func calcularImpuesto(bruto float64) float64 {
	var descuento float64 = 0

	if bruto > 50000 {
		descuento = bruto * 0.17

		if bruto > 150000 {
			adicional := bruto * 0.10
			descuento = descuento + adicional
		}
	}

	return descuento
}
