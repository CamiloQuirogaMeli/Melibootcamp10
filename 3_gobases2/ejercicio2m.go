package main

// Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y nos
// devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
func avgStudent(scores ...int) float64 {
	var accum float64
	for _, score := range scores {
		accum += float64(score)
	}
	return accum / float64(len(scores))
}
