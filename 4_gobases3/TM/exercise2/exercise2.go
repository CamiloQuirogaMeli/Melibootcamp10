package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(fileName string) {
	f, err := os.Open("./products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("ID\t\t\t\t Precio\t\t Cantidad")
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		fmt.Printf("%s\t\t%s\t\t\t\t%s\n", line[0], line[1], line[2])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	readFile("../products.txt")
}
