package main

import (
	"fmt"
)

func main() {
	num := 14
	month, err := intToMonth(num)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(month)
}
