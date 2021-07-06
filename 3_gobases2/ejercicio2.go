package main

import "errors"

func promedioCalificaciones(calificaciones ...float64) (float64, error) {

	var suma float64 = 0

	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("una calificacion es negativa")
		}
		suma += calificacion
	}

	var promedio float64 = (suma / float64(len(calificaciones)))

	return promedio, nil

}
