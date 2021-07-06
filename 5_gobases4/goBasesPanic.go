package main

import (
	"context"
	"fmt"
	"time"
	//"os"
)

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, " hace woof woof")
}
func panico() {
	/* RECEPTORES NULOS
	s := &Dog{"Sammy"}
	s = nil
	s.WoofWoof()
	*/

	/* INDEX OUT OF BOUNDS
	Milanesa := []string{
		"pollo",
		"carne",
		"pescado",
	}
	fmt.Println("la mas rica es: ", Milanesa[len(Milanesa)])
	*/

	/* EJEMPLO DE PANIC POR NO ENCONTRAR UN ARCHIVO
	fmt.Println("Iniciando... ")
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Fin")
	*/

	defer func() { //func(tipo funcion)
		fmt.Println("Funcion que se ejecuta por mas que haya un panic")
	}() //para que defer la ejecute

	panic("Panico escenico!!!") //si etuviera antes del defer aun asi esa funcion no se ejecuta

}

func main() {
	num := 3
	isPair(num)
	fmt.Println("Ejecución completada!")

	contextos()
	panico()
}
func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if (num % 2) != 0 {
		panic("no es un número par")
	}
	fmt.Println(num, " es un número par!")
}

func contextos() {

	ctx0 := context.Background()
	ctx0 = context.WithValue(ctx0, "saludo", "hola digital house!! Quiero una milanesa")
	saludoWrapper(ctx0)

	ctx := context.Background()
	deadline := time.Now().Add(time.Second * 5)
	ctx, _ = context.WithDeadline(ctx, deadline)
	<-ctx.Done()
	fmt.Println(ctx.Err().Error())

	ctx2 := context.Background()
	ctx2, _ = context.WithTimeout(ctx2, time.Second*5)
	<-ctx2.Done()
	fmt.Println(ctx2.Err().Error())
}

func saludoWrapper(ctx context.Context) {
	saludo(ctx)
}
func saludo(ctx context.Context) {
	fmt.Println(ctx.Value("saludo"))
}
