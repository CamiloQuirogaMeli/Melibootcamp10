package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	dat, err := ioutil.ReadFile("../exercise1/productList.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		plainText := string(dat)
		products := formatData(plainText)
		printFormatedData(products)
	}
}

type Product struct {
	id     int
	price  float64
	amount int
}

func (p Product) toString() string {
	return fmt.Sprint(p.id, ";", p.price, ";", p.amount, "\n")
}

func formatData(data string) []Product {
	var formatedData []Product
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		tokens := strings.Split(line, ";")
		if len(tokens) != 0 {
			product, err := readProduct(tokens)
			if err == nil {
				formatedData = append(formatedData, product)
			}
		}
	}
	return formatedData
}

func readProduct(args []string) (Product, error) {

	product := Product{}
	theId, err := strconv.Atoi(args[0])
	if err != nil {
		return product, err
	} else {
		product.id = theId
	}

	theAmount, err := strconv.Atoi(args[2])
	if err != nil {
		return product, err
	} else {
		product.amount = theAmount
	}
	thePrice, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return product, err
	} else {
		product.price = thePrice
	}

	return product, nil
}

func printFormatedData(products []Product) {
	fmt.Println("ID\t\tPrice\tAmount")
	acumPrice := 0
	for _, product := range products {
		fmt.Print(product.id, "\t\t", product.price, "\t", product.amount, "\n")
		acumPrice += int(product.price) * product.amount
	}

	fmt.Print("\t\t", acumPrice, "\t\n")
}
