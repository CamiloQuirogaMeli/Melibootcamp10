package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		fmt.Println("La palabra ingresada tiene", len(text), " letras")

		for i := 0; i < len(text); i++ {
			fmt.Println("\"", string((text[i])), "\"")

		}
	}
}
