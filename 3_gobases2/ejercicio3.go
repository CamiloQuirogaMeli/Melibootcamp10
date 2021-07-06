package main

func calcularSalario(minutosPorMes float64, categoria string) (salario float64) {

	var horasPorMes float64 = minutosPorMes / 60

	switch categoria {
	case "C":
		return (horasPorMes * 1000)
	case "B":
		return (horasPorMes * 1500) * 1.2
	case "A":
		return (horasPorMes * 3000) * 1.5
	}
	return 0
}
