package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Product struct {
	Id 			int64
	Price 		float64
	Quantity 	int64
}

func ParseProduct(s string) (p Product, err error) {
	fields := strings.Split(s, ";")

	p.Id, err = strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return
	}

	p.Price, err = strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return
	}
	
	p.Quantity, err = strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return
	}

	return
}

func main() {

	data, err := ioutil.ReadFile("products.txt")
	
	if err == nil {
		
		fmt.Println("ID\t\t    Precio   Cantidad")
		
		lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
		var total float64 = 0
		for _, line := range lines {

			product, err := ParseProduct(line)
			if err == nil {
				total += product.Price * float64(product.Quantity)
			}

			fmt.Printf("%d\t\t%10.2f%11d\n", product.Id, product.Price, product.Quantity)
		}

		fmt.Printf("\t\t%10.2f\n",total)
	}
}