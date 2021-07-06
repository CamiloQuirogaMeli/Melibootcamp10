package main

import "fmt"

func main() {
	readWord("paralelepipedo")
}

func readWord(str string) {
	var n = len(str)

	fmt.Println(n, str)

	for i := 0; i < n; i++ {
		fmt.Println(string(str[i]))
	}
}
