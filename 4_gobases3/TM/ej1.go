package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	d1 := []byte(" id:1 \n precio:10 \n cantidad:5 \n\n id:1 \n precio:10 \n cantidad:5 \n")
	err := ioutil.WriteFile("./data", d1, 0644)
	fmt.Println(err)

}
