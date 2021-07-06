package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	escribir := []byte("1000, 30000, 2; 1001, 15000, 1; 1002, 500, 20;")
	err := ioutil.WriteFile("./archivoEjer", escribir, 0644)
	fmt.Println(err)
}
