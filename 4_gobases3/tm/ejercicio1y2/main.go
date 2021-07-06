package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileDir = "./products.txt"

func listProducts(f *os.File) {
	rd := bufio.NewReader(f)
	fmt.Printf("ID \t Price \t Amount\n")
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		data := strings.Split(string(line), ";")
		fmt.Printf("%s \t %s \t %s\n", data[0], data[1], data[2])
	}
}

func saveProduct(f *os.File) {
	var ID uint64
	var price float64
	var amount int
	fmt.Print("Enter ID -> ")
	fmt.Scanln(&ID)
	fmt.Print("Enter price -> ")
	fmt.Scanln(&price)
	fmt.Print("Enter amount -> ")
	fmt.Scanln(&amount)
	in := fmt.Sprintf("%d;%.2f;%d\n", ID, price, amount)
	w := bufio.NewWriter(f)
	_, err := w.WriteString(in + "\n")
	if err != nil {
		panic(err)
	}
	w.Flush()
	fmt.Println("Product added")
}

func main() {
	f, err := os.OpenFile(fileDir, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	listProducts(f)
	var in string
	fmt.Print("Do you want to add a product? (y/n) \n")
	fmt.Scanln(&in)
	if in == "n" {
		os.Exit(0)
	}
	saveProduct(f)
}
