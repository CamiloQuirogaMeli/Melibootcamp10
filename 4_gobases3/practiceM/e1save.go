package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type Product struct {
	Id    string
	Price float64
	Cant  int
}

func main() {
	file := createFile()
	defer closeFile(file)
	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("A32_%d", i)
		price := 10.0
		cant := 2
		writeFile(file, id, price, cant)
	}
	showFileContent(file)
}

func createFile() *os.File {
	filename := "inventory.csv"
	var file *os.File
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("no Existe el archivo, creandolo...")
		f, e := os.Create(filename) //os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if e != nil {
			panic(e)
		} else {
			file = f
			fmt.Fprintln(f, "ID,Precio,Cantidad")
			fmt.Println("archivo creado correctamente")
		}
	} else {
		fmt.Println("el archivo ya existe")
		f, e := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if e != nil {
			panic(e)
		} else {
			file = f
		}

	}
	return file
}

func writeFile(file *os.File, id string, price float64, cant int) {
	fmt.Println("escribiendo...")
	fmt.Fprintln(file, fmt.Sprintf("%s,%.2f,%d", id, price, cant))
}

func closeFile(file *os.File) {
	fmt.Println("cerrando archivo...")
	file.Close()
}

func showFileContent(file *os.File) {
	fmt.Println("FGGG")
	csvFile, err := os.Open("inventory.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.AlignRight)
	auxPrice := 0.0
	auxCant := 0
	auxSubTot := 0.0
	for key, line := range csvLines {
		if key == 0 {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t", line[0], line[1], line[2]))
			w.Flush()
		} else {
			product := Product{
				Id:    line[0],
				Price: 0,
				Cant:  0,
			}
			if f, err := strconv.ParseFloat(line[1], 64); err == nil {
				product.Price = f
				auxPrice = f
			}
			if i, err := strconv.Atoi(line[2]); err == nil {
				product.Cant = i
				auxCant = i
			}
			auxSubTot += auxPrice * float64(auxCant)
			fmt.Fprintln(w, fmt.Sprintf("%s\t%.2f\t%d\t", product.Id, product.Price, product.Cant))
			w.Flush()
		}
	}
	fmt.Fprintln(w, fmt.Sprintf("\t%.2f\t\t", auxSubTot))
	w.Flush()
	fmt.Println()
}
