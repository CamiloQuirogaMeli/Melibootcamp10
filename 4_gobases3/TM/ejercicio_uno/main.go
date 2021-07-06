package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile(content string, path string) {
	d1 := []byte(content)
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		fmt.Println("error")
	}
}

func main() {

	productsList := "[{\"id\": 1, \"precio\": 456, \"cantidad\": 6}, {\"id\": 2, \"precio\": 234, \"cantidad\": 2}, {\"id\": 3, \"precio\": 23, \"cantidad\": 12}]"
	writeFile(productsList, "./productos.json")
}
