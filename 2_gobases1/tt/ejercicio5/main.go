package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileDir = "./alumnos.txt"

// List all students saved in the file given
func listStudents(f *os.File) {
	rd := bufio.NewReader(f)
	fmt.Println("Students list:")
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(line))
	}
}

// saveStudent save a student in the file given
func saveStudent(f *os.File) {
	var in string
	fmt.Print("Enter name -> ")
	fmt.Scanln(&in)
	w := bufio.NewWriter(f)
	_, err := w.WriteString(in + "\n")
	if err != nil {
		panic(err)
	}
	w.Flush()
	fmt.Println("Added student")
}

func main() {
	f, err := os.OpenFile(fileDir, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	listStudents(f)
	var in string
	fmt.Print("Do you want to add another? (y/n) \n")
	fmt.Scanln(&in)
	if in == "n" {
		os.Exit(0)
	}
	saveStudent(f)
}
