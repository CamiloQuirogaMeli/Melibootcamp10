package main

import "fmt"

//La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla, para ello es necesario crear una
//aplicación tener una variable con la palabra e imprimir la cantidad de letras que tiene la misma, luego imprimí cada una de las letras.

func main()  {
	var word = "PALABRA"
	fmt.Println("Palabra: " + word)
	fmt.Println("Cantidad de letras: ", len(word))
	fmt.Println("Letras: ")
	for i := 0; i < len(word); i++ {
		j := i + 1
		fmt.Println(word[i:j])
	}
}