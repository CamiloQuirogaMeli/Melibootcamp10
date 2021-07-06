package main

import (
	"fmt"
)

func mean(nums []int) float64 {
	var size int = 0
	var sum int = 0
	for _, num := range nums {
		if num < 0 {
			return float64(-1)
		}
		sum += num
		size++
	}
	return float64(sum) / float64(size)
}

func main() {
	var grades []int
	var size int
	fmt.Print("Escriba la cantidad de notas: ")
	fmt.Scanln(&size)
	for i := 0; i < size; i++ {
		var temp int
		fmt.Print("Escriba la nota #", i+1, ": ")
		fmt.Scanln(&temp)
		if temp < 0 {
			fmt.Println("Error, se insertó un número negativo")
			i--
		} else {
			grades = append(grades, temp)
		}
	}
	var rmean float64 = mean(grades)
	fmt.Println("Promedio: ", rmean)
}
