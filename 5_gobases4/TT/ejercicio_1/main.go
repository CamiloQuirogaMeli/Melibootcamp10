package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("customers.txt")
	
	defer func (){fmt.Println("ejecucion finalizada")}()
	if err != nil{
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Printf("Contenido del archivo: %s\n", string(content))
	}
}
   