package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	i1 := Item{id: 1, price: 2500.00, quantity: 5}
	i2 := Item{id: 2, price: 1500.00, quantity: 10}
	i3 := Item{id: 3, price: 500.00, quantity: 2}

	data := []byte("id, price, quantity\n" + i1.getItem() + ";\n" + i2.getItem() + ";\n" + i3.getItem() + ";")

	err := ioutil.WriteFile("./dat1", data, 0644)

	if err != nil {
		fmt.Println(err)
	}
}

type Item struct {
	id       uint
	price    float64
	quantity int
}

func (i Item) getItem() string {
	return fmt.Sprintf("id: %d, price: %f, quantity: %d", i.id, i.price, i.quantity)
}
