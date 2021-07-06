package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Prices struct {
	Prices []Price
}

type Price struct {
	ID       int     `json:"id"`
	PRECIO   float32 `json:"precio"`
	CANTIDAD int     `json:"cantidad"`
}

func main() {

	jsonFile, err := os.Open("./productos.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var prices Prices
	json.Unmarshal(byteValue, &prices)

	fmt.Printf("id\t precio\t\t\t\t cantidad\n")
	for i := 0; i < len(prices.Prices); i++ {
		fmt.Printf("%d \t %.2f \t\t\t %d \n", prices.Prices[i].ID, prices.Prices[i].PRECIO, prices.Prices[i].CANTIDAD)
	}

}
