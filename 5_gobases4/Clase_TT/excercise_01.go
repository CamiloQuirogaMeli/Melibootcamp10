package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile()
	fmt.Println("Execution completed")
}

func readFile() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := ioutil.ReadFile("./customers.txt")

	if err != nil {
		panic("the indicated file was not found or is corrupted")
	}
}
