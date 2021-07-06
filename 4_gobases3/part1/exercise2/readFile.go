package main

import (
	"bufio"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total float64
	dataFile, _ := os.Open("./../exercise1/archivo.txt")
	f.Println("ID\t\t\tPrecio\t\tCantidad")
	scanner := bufio.NewScanner(dataFile)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ";")
		v1, _ := strconv.ParseFloat(values[2], 64)
		v2, _ := strconv.ParseFloat(values[3], 64)
		total += v1 * v2
		f.Printf("%s\t\t\t%s\t\t%s\n", values[0], values[1], values[2])

	}
	f.Print("\t\t\t------\t\t------\n")
	f.Printf("\t\t\tTotal:\t\t%.2f\n", total)
}
