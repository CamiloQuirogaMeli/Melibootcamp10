package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile() {
	data, err := ioutil.ReadFile("../archivo.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\t\t%7s\t\t%8s\n", "ID", "PRECIO", "CANTIDAD")

		var totalPrice float64 = 0
		lines := strings.Split(string(data), ";")
		for _, line := range lines {
			splitLine := strings.Split(string(line), " ")
			fmt.Printf("%s\t\t%7s\t\t%8s\n", splitLine[0], splitLine[1], splitLine[2])
			//fmt.Println(splitLine[0], "\t", splitLine[1], "\t", splitLine[2])
			priceToFloat, err := strconv.ParseFloat(splitLine[1], 64)
			if err != nil {
				fmt.Println(err)
			} else {
				cantToFloat, err := strconv.ParseFloat(splitLine[2], 64)
				if err != nil {
					fmt.Println(err)
				} else {
					totalPrice += priceToFloat * cantToFloat
				}
			}
		}
		fmt.Println("\t\t", totalPrice)
	}
}

func main() {
	readFile()
}
