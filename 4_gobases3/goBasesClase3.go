package main

import (
	"fmt"
	//"os"
	//"io/ioutil"
	"time"
)

func main() {
	//PAQUETE FMT
	//verbos de impresion --> indican como imprimir una variable
	/*
		//.Print()
		const NOMBRE, EDAD = "Kim", 22
		fmt.Print(NOMBRE, " tiene ", EDAD, " años.\n ")

		//en .Println() no hace falta al final el \n
		//.Printf()
		fmt.Printf("%s tiene %d años\n", NOMBRE, EDAD)

		//.Sprint()
		//tenemos que ingresar los verbos %s o %d
		res := fmt.Sprintf("%s tiene %d años\n", NOMBRE, EDAD)
		fmt.Print(res)
	*/
	//PAQUETE OS
	/*
		//.Getenv()
		valor := os.Getenv("HOME")
		fmt.Println(valor)

		//.LookupEnv() --> verifica si existe una variable de entorno
		valor, existe := os.LookupEnv("HOME")
		fmt.Println(valor, existe)
		//si existe pero vacia es true, sino es false

		//.Setenv()
		err := os.Setenv("MI_VARIABLE", "mi valor")
		fmt.Println(err)
	*/

	//PAQUETE IO/IOUTIL
	/*
		//.ReadFile() --> para leer archivo
		dat, err := ioutil.ReadFile("./miArchivo.txt")
		fmt.Println(dat, err)

		//.WriteFile() --> para escribir archivo
		d1 := []byte("hello\ngo\n")
		//err := ioutil.WriteFile("./dat1", d1, 0644)
		fmt.Println(d1)
	*/

	//CLASE TT
	//punteros
	punteros()
	valor := 9
	Incrementar(&valor)
	fmt.Println("El valor cambio a :", valor)

	//go routines
	/*
		for i := 0; i < 10; i++ {
			go proceso(i) //la sentencia GO garantiza la concurrencia, no el paralelismo
		}
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("Terminó el programa")
	*/

	//canales
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Terminó el programa", <-c)
	}

}

func punteros() {

	v := 19
	h := 20
	var p1 *int
	//var p2 = new(int)
	p2 := &h
	p3 := &v
	fmt.Println("v: \n", v)
	*p3 = 42
	fmt.Println("v: \n", v)
	// %T nos permite imprimir el tipo de dato de la variable
	fmt.Println("p2:  \n", p1) //nos da la posicion en memoria del puntero
	fmt.Print(*p2)
	fmt.Printf("\np2: %T \n", p2)
	fmt.Printf("p3: %T \n", p3)

}

func Incrementar(v *int) {
	(*v)++
}

func proceso(i int, c chan int) {
	fmt.Println(i, "-Inicio")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Fin")
	c <- i //envio valor a canal
}
