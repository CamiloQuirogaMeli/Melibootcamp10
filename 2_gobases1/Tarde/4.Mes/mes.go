package main

import "fmt"

/* Se podria hacer con muchos else if, con el condicional switch, con un map e incluso podria existir una libreria
que ya lo haga. En este caso escogeré con un map porque me permite asignarle a una clave, un valor */
func main() {
	mes := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio",
		8: "Agosto", 9: "Septiempre", 10: "Octubre", 11: "Noviempre", 12: "Diciempre"}
	fmt.Println(mes[7])
}
