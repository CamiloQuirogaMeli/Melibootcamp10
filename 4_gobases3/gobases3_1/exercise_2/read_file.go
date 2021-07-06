package main

import (
	f "fmt"
	"io/ioutil"
	strco "strconv"
	str "strings"
)

func main() {
	data, err := ioutil.ReadFile("gobases3_1/save_file.txt")
	if err != nil {
		f.Println("Error while reading the file: ", err)
	}
	printTable(data)
}

func printTable(data []byte) {
	var total float64
	products := str.Split(string(data), ";")
	f.Println("\nID \t Price \t\t Amount")
	for _, product := range products {
		data := str.Split(product, ",")
		price, _ := strco.ParseFloat(data[1], 64)
		amount, _ := strco.ParseFloat(data[2], 64)
		total += price * amount
		f.Println(data[0], "\t", data[1], "\t\t", data[2])
	}
	f.Print("\n\t ", total, "\n")
}
