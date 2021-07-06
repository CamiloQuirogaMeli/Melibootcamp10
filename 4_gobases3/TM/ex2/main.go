package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	var total float64
	dataFile, err := os.Open("../data.csv")
	if err != nil {
		fmt.Println("File not found")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	fmt.Fprintln(w, "ID\t\t\tPrecio\t\tCantidad")
	scanner := bufio.NewScanner(dataFile)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ";")
		id, err := strconv.Atoi(values[0])
		price, err := strconv.ParseFloat(values[1], 64)
		inv, err := strconv.Atoi(values[2])
		if err != nil {
			fmt.Println("File corrupted")
			return
		}
		total += price * float64(inv)
		fmt.Fprintln(w, fmt.Sprintf("%d\t\t\t%.2f\t\t%d", id, price, inv))

	}
	fmt.Fprintln(w, fmt.Sprintf("\t\t\t------\t\t------"))
	fmt.Fprintln(w, fmt.Sprintf("Total:\t\t\t%.2f", total))
	w.Flush()
}
