package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Productos struct {
	Id       string
	Precio   float64
	Cantidad int
}

func main() {
	content, err := ioutil.ReadFile("../Ejercicio1/Pruebita")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(content))

	// fmt.Println("ID\t\t\t\t Precio\t\t Cantidad")
	// for i, value := range texto {
	// 	fmt.Println(value["Id"], "t\t\t\t", value[i]["Precio"], "t\t", value[i]["Cantidad"])
	// }

}
