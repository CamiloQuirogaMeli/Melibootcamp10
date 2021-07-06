package main

import (
	"fmt"
	"io/ioutil"
)

type Product struct {
	Id string
	Price float64
	Amount int
}

func main(){
	product := Product{
		Id: "1",
		Price: 15.25,
		Amount: 8,
	}
	product2 := Product{
		Id: "2",
		Price: 13.66,
		Amount: 3,
	}
	products := []Product{product, product2}
	var str string
	for _, p := range products {
		line := fmt.Sprintf("%s;%.2f;%d\n", p.Id, p.Price, p.Amount)
		str += line
	}
	data := []byte(str)
	ioutil.WriteFile("./archivo.txt", data, 0644)

}
