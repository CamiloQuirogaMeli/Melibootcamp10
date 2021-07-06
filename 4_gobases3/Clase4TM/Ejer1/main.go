package main

import (
	"fmt"
	"io/ioutil"
)

type product struct {
	id     int64
	price  float64
	amount int64
}

func (p product) Info() string {
	return fmt.Sprintf("%d;%.2f;%d;", p.id, p.price, p.amount)

}

func main() {
	filename := "./miArchivo.csv"
	p1 := product{1, 1234, 3}
	p2 := product{2, 7777, 2}
	p3 := product{3, 4321, 1}
	data := ""
	data += p1.Info()
	data += p2.Info()
	data += p3.Info()
	databyte := []byte(data)
	ioutil.WriteFile(filename, databyte, 0644)
}
