package main

import "fmt"

type Matrix struct{
	Valores []float64
	Alto int
	Ancho int
}

func (m Matrix) Cuadratica() bool{
	if m.Alto == m.Ancho{
		return true
	}else{
		return false
	}
}

func (m Matrix) Maximo(){
	if len(m.Valores) == 0{
		fmt.Println("La matriz no posee valores para calcular el maximo.")
	}else {
		var max float64 = 0
		for i := 0; i < len(m.Valores); i++ {
			if max != 0 {
				max = m.Valores[i]
			}
			if max < m.Valores[i]{
				max = m.Valores[i]
			}
		}
		fmt.Println("El valor maximo de la matriz es: ", max)
	}

}

func (m Matrix) Set() {
	if len(m.Valores) != m.Ancho*m.Alto{
		fmt.Println("La cantidad de valores no corresponde a las dimensiones de la matriz.")
	}
}

func (m Matrix) Print(){
	fmt.Println("Alto = ", m.Alto)
	fmt.Println("Ancho = ", m.Ancho)

	if len(m.Valores) == 0{
		fmt.Println("La matriz esta vacia.")
	}
	for i := 0; i < m.Alto; i++ {
		fmt.Println(m.Valores[i * m.Ancho : i * m.Ancho + m.Alto])

	}
}

func main(){
	mat := Matrix{
		Valores: []float64{3,6,2,6,7.3,10,8,3.4,7.2,4,32,9.8},
		Alto: 4,
		Ancho: 4,
	}

	Matrix.Set(mat)
	mat.Maximo()
	mat.Print()

}

/*Una empresa de inteligencia artificial necesita tener una funcionalidad para crear matrices en
Go.
Para ello requieren una estructura Matrix que tenga los métodos:
- Set: Recibe una matrix de flotantes e inicializa los valores en la estructura Matrix
- Print: Imprime por pantalla la matrix de una formas más visible (Con los saltos de línea
entre filas)
La estructura matrix debe guardar la matrix, el tamaño de alto, el tamaño de ancho, si es
cuadrática y el valor máximo.*/
