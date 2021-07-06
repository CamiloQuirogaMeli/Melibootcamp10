package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	defer func() {
		fmt.Println("Execution finished")
	}()

	data, err := ioutil.ReadFile("./file.txt")

	if err != nil {
		panic("the indicated file was not found or is corrupted")
	} else {
		fmt.Println(string(data))
	}

}
