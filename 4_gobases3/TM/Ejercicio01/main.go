package main

import (
	"io/ioutil"
)

func main() {
	createFile()
}

func createFile() {
	fileValues := "1234 70.50 20;25768 89.00 5;3980 50.75 50;86950 90.00 5;7890 80.2 10.00"
	textFile := []byte(fileValues)
	ioutil.WriteFile("../archivo.txt", textFile, 0644)
}
