package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name      string
	Price     int
	Published bool
}

func main() {

	// p := product{
	// 	Name:      "MacBook Pro",
	// 	Price:     1500,
	// 	Published: true,
	// }
	// jsonData, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(jsonData))

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`
	p := product{}
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}
