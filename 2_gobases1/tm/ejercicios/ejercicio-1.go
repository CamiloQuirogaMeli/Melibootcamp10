package ejercicios

import "fmt"

func PrimerEjercicio() {
	// Declaramos la variable con la palabra
	var word = "Palabra"

	// Imprimimos la cantidad de letras que tiene la palabra
	fmt.Println(len(word))

	// Imprimimos cada una de las letras de la palabra
	for _, char := range word {
		fmt.Println(string(char))
	}
}
