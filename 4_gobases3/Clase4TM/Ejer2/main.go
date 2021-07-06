package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type product struct {
	id     int64
	price  float64
	amount int64
}

func (p product) Info() string {
	return fmt.Sprintf("id:%d;price:%.2f;amount:%d;", p.id, p.price, p.amount)

}

func main() {
	dat, _ := ioutil.ReadFile("miArchivo.csv")
	values := strings.Split(string(dat), ";")
	fmt.Println("ID \tPrecio\t Cantidad")
	total := 0.0
	for i := 0; i < len(values)-1; i += 3 {
		index, _ := strconv.Atoi(values[i])
		price, _ := strconv.ParseFloat(values[i+1], 64)
		amount, _ := strconv.Atoi(values[i+2])
		total += price * float64(amount)
		fmt.Printf("%d \t%.2f\t\t%d\n", index, price, amount)
	}
	fmt.Printf(" \t%.2f\t\n", total)
}
