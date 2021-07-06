package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	content, err := ioutil.ReadFile("customers.txt")

	defer func() { fmt.Println("execution completed") }()

	if err != nil {
		panic("the indicated file was not found or is corrupted")
	} else {
		fmt.Println(string(content))
	}

}
