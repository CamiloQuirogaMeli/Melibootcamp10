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

	p1 := Productos{Id: "asdasdasd", Precio: 7854.8, Cantidad: 150}
	p2 := Productos{Id: "123123", Precio: 8458, Cantidad: 210}
	p3 := Productos{Id: "1234hdy", Precio: 8541.9, Cantidad: 77}
	sliceP := []Productos{p1, p2, p3}

	// b1, err := json.Marshal(p1)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }

	// b2, err := json.Marshal(p2)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }

	// b3, err := json.Marshal(p3)
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }

	fmt.Println("Ingresá el titulo del archivo: ")
	var titulo string
	fmt.Scanln(&titulo)

	txt := []byte(fmt.Sprint(sliceP))
	// txt1 := []byte(fmt.Sprint(p2))
	// txt2 := []byte(fmt.Sprint(p3))

	err := ioutil.WriteFile(titulo, txt, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// file, err := os.OpenFile(titulo, os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer file.Close()
	// if _, err := file.WriteString(string(b2)); err != nil {
	// 	log.Fatal(err)
	// }

	// file, err3 := os.OpenFile(titulo, os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err3)
	// }
	// defer file.Close()
	// if _, err := file.WriteString(string(b3)); err != nil {
	// 	log.Fatal(err)
	// }

}
