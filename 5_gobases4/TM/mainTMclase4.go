package main

/*
Ejercicio 1
1. En tu función “main”, define una variable llamada “salary” y asignale un valor de tipo
“int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
“salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe
pagar impuesto”.

Ejercicio 2
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”.

EJercicio 3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de
error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y
el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

//"errors"
//"fmt"

func main() {
	/*
		//ejercicio 1
		fmt.Println("Error ejercicio 1")
		var salary int = 2000
		_, err := myCustomErrorTest(salary)
		if err != nil {
			fmt.Println(err)
		}

		//ejercicio 2
		fmt.Println("Error ejercicio 2")
		salary = 2000
		if salary < 15000 {
			fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		} else {
			fmt.Println("Debe pagar impuesto")
		}

		//ejercicio 3
		fmt.Println("Error ejercicio 3")
		salary = 15000
		if salary < 150000 {
			eror := fmt.Errorf("error: el minimo imponible es de 150000 y el salario ingresado es de: ", salary)
			fmt.Println("error ocurrido: \n", eror)
		} else {
			fmt.Println("Salario supera mínimo imponible")
		}
	*/

}

/*
type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(salario int) (int, error) { //devolvemos un objeto de la interface error, que está siendo creado en el momento
	if salario < 15000 {
		return 400, &myCustomError{ //el & es para usarlo fuera de la funcion y manipular el objeto
			status: salario,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		fmt.Println("Debe pagar impuesto")
	}
	return 200, nil
}
*/
