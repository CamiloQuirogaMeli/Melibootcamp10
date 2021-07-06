package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//los atributos si o si deben ir en Mayuscula para exportarlos
//y su posterior transformacion en json
type Product struct {
	Name      string
	Price     float64
	Published bool
}

func main() {
	p := Product{
		Name:      "Macbook Pro",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	jsonDataUN := `{"Name": "MacBook Air", "Price": 900, "Published": false}`

	p1 := Product{} //si no se pasa como puntero se estaría duplicando

	if err2 := json.Unmarshal([]byte(jsonDataUN), &p1); err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(p1)

}
