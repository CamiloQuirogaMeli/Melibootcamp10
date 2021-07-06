package main

import "fmt"

func gestorOperaciones(operacion string) func(values...int) float32{

	switch operacion{
	case "minimo":
		return calculoMin

	case "maximo":
		return calculoMax

	case "promedio":
		return calculoProm
	}
	return nil
}

func calculoMin (nros...int) float32{

	var min int
	var flag bool = false

	for _, element := range nros {
		if flag == false{
			min = element
			flag = true
		}else if element < min{
			min = element
		}
	}
	return float32(min)
}

func calculoMax (nros...int) float32{
	var max int
	var flag bool = false

	for _, element := range nros {
		if flag == false{
			max = element
			flag = true
		}else if element > max{
			max = element
		}
	}
	return float32(max)
}

func calculoProm (nros...int) float32{
	var prom float32 = 0.0
	var acum int

	for _, element := range nros {
		acum += element
	}

	prom = float32(acum/len(nros))

	return prom
}

func main(){
	gestor := gestorOperaciones("promedio")

	proceso := gestor(5, 8, 2, 5, 8, 9)

	fmt.Println(proceso)
}


/*Una universidad de Entre Ríos necesita sacar estadísticas de los alumnos de cada curso,
requiriendo calcular el mínimo, máximo y promedio.

Se solicita generar una función que le indique qué tipo de cálculo queramos realizar (mínimo,
máximo o promedio) y nos devuelve otra función ( y un error en caso que el cálculo no esté
definido) que le podamos pasar una cantidad N de enteros y nos devuelva el cálculo que
indicamos en la función anterior*/