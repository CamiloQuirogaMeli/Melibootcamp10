package main


import (
"encoding/json"
"fmt"
"log"
)

type product struct {
	Name      string `json:"name,omitempty"`
	Price     int    `json:"price,omitempty"`
	Published bool   `json:"published,omitempty"`
}

func main() {
	jsonData := `{"name": "MacBook Air", "price": 900, "published": true}`

	p := product{}

	err := json.Unmarshal([]byte(jsonData), &p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}