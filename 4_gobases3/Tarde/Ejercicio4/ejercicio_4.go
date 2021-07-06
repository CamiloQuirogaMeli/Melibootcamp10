package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insercion(items []int, c chan []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	c <- items
}

func seleccion(items []int, c chan []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	c <- items
}

func burbujeo(items []int, c chan []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	c <- items
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(100)
	variable3 := rand.Perm(100)

	/*variable4 := rand.Perm(1000)
	variable5 := rand.Perm(1000)
	variable6 := rand.Perm(1000)*/

	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	/*c4 := make(chan []int)
	c5 := make(chan []int)
	c6 := make(chan []int)*/

	inicio100 := time.Now()

	go insercion(variable1, c1)
	go seleccion(variable2, c2)
	go burbujeo(variable3, c3)

	<-c1
	<-c2
	<-c3

	diff100 := time.Now().Sub(inicio100) //resto la hora de cuando arrancan las rutinas y la hora de cuando obtuve el resultado de los 3 canales
	fmt.Println(diff100)

}
