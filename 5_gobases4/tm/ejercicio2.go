package main

import (
	"errors"
	"fmt"
)

func main()	{
	salary := 140000
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mimimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}

