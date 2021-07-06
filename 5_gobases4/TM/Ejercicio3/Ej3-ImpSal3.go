package main

import (
	"fmt"
)


func main(){
	var salary int = 160000

	if salary < 150000{
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println("problema - ", err)
	}else{
		fmt.Println("Debe pagar impuesto.")
	}

}


/* EJERCICIO1
1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”.*/

/* EJERCICIO2
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”.*/

/* EJERCICIO3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
 */