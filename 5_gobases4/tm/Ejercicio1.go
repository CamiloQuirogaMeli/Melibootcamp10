package main

import "fmt"

type myError struct {
	message string
}

func (err myError) Error() string {
	return err.message
}

func main()	{

	salary := 160000
	status := errorTest(salary)
	if status != nil {
		fmt.Println(status)
	} else {
		fmt.Println("Debe pagar impuesto")
	}


}

func errorTest(salary int) error{
	if salary <150000 {
		return myError{
			message: "error: el salario ingresado no alcanza el mimimo imponible",
		}
	}
	return nil
}
