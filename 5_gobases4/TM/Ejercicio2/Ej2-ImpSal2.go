package main

import (
	"errors"
	"fmt"
)


func main(){
	var salary int = 149999

	if salary < 150000{
		fmt.Println(errors.New("error: el salario ingresado no alcanza el minimo imponible"))
	}else{
		fmt.Println("Debe pagar impuesto")
	}

}


/* EJERCICIO1
1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”.*/

/* EJERCICIO2s
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”.*/
