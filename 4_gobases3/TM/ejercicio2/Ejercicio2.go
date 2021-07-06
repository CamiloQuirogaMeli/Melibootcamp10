package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.Open("../ejercicio1/archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo")
	} else {
		fmt.Println("ID\t\t\tPrecio\tCantidad")
		scanner := bufio.NewScanner(file)
		total := 0.0
		for scanner.Scan() {
			line := scanner.Text()

			values := strings.Split(line, ";")
			v1, _ := strconv.ParseFloat(values[1], 64)
			v2, _ := strconv.ParseFloat(values[2], 64)
			total +=  v1 * v2
			fmt.Printf("%s\t\t\t%s\t%s\n", values[0], values[1], values[2])

		}
		fmt.Printf("\t\t\t%.2f\n", total)
	}
}
