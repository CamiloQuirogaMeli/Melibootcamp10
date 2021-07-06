package main

import (
	"fmt"
	"io/ioutil"
)

type Sale struct {
	ProductId int
	Amount    int
	Price     float64
}

func main() {
	sales := []Sale{{1, 2, 350},
		{2, 2, 350},
		{8, 21, 350},
		{11, 23, 350},
		{15, 20, 350},
		{1, 1, 350}}

	data := ""
	for _, sale := range sales {
		line := ""
		line += fmt.Sprint(sale.ProductId) + " "
		line += fmt.Sprint(sale.Amount) + " "
		line += fmt.Sprint(sale.Price) + ";\n"
		data += line
	}

	err := ioutil.WriteFile("ventas.dat", []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
	}

}
