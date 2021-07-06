package main

import "fmt"
import "errors"

//Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso, requiriendo calcular el mínimo, máximo y promedio.

//Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo, máximo o promedio) y nos devuelve otra función (y un error en caso que el cálculo
//no esté definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que indicamos en la función anterior
const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

func main()  {
	fmt.Println("Se va a calcular las estadísticas para sólo un estudiante...")
	minFunc, err1 := operacion(minimo)
	promFunc, err2 := operacion(promedio)
	maxFunc, err3 := operacion(maximo)
	_, err4 := operacion("na")

	if(err1 == nil){
		valorMinimo := minFunc(2,3,3,4,1,2,4,5)
		fmt.Println("La nota mínima del estudiante es:", valorMinimo)
	} else {
		fmt.Println("ERROR:", err1)
	}

	if(err3 == nil){
		valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
		fmt.Println("La nota máxima del estudiante es:", valorMaximo)
	} else {
		fmt.Println("ERROR:", err3)
	}

	if(err2 == nil){
		valorPromedio := promFunc(2,3,3,4,1,2,4,5)
		fmt.Println("El promedio del estudiante es:", valorPromedio)
	} else {
		fmt.Println("ERROR:", err2)
	}

	if(err4 == nil){
		fmt.Println(nil)
	} else {
		fmt.Println("ERROR:", err4)
	}
	
	
	

}

func operacion(op string) (func(notas ...int) float64, error) {
	switch op {
	case minimo:
		return min, nil
	case promedio:
		return prom, nil
	case maximo:
		return max, nil
	}

	return nil, errors.New("El cálculo que desea hacer no está definido")
}

func min(notas ...int) float64 {
	min := 10

	for _, value := range notas {
		if value < min {
			min = value
		}
	}

	return float64(min)
}

func max(notas ...int) float64 {
	max := 0

	for _, value := range notas {
		if value > max {
			max = value
		}
	}

	return float64(max)
}

func prom(notas ...int) float64 {
	var prom float64

	for _, value := range notas {
		prom += float64(value)
	}

	return prom / float64(len(notas))
}