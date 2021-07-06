package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	producto := []byte("12345\t500\t2")
	err := ioutil.WriteFile("./product.txt", producto, 0644)
	err = ioutil.WriteFile("./product.txt", producto, 0644)
	err = ioutil.WriteFile("./product.txt", producto, 0644)

	fmt.Println(err)
}
