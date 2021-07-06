package main

import(
	"fmt"
	"os"
	)

type miErrorSalary struct{

}

type error interface {
	Error() string
}

func (e *miErrorSalary) Error() string{
	return "error: el salario ingresado no alcanza el minimo imponible"
}

func errorTest (salario int) (string, error){
	if salario < 150000{
		return "", &miErrorSalary{}
	}else{
		return "Debe pagar impuesto", nil
	}
}

func main(){
	var salary int = 170000

	impString, err := errorTest(salary)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(impString)

}


/*1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”.*/
