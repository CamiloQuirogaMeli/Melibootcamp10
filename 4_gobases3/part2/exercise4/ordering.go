package main

import (
	f "fmt"
	"math/rand"
	"time"
)

func Insertion(Array []int, ch chan time.Duration) {
	start := time.Now()
	var auxiliar int
	for i := 1; i < len(Array); i++ {
		auxiliar = Array[i]
		for j := i - 1; j >= 0 && Array[j] > auxiliar; j-- {
			Array[j+1] = Array[j]
			Array[j] = auxiliar
		}
	}
	ch <- time.Since(start)
}

func Bubble(ListaDesordenada []int, ch chan time.Duration) {
	start := time.Now()
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	ch <- time.Since(start)
}

func Selection(ListaDesordenada *[]int, ch chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(*ListaDesordenada); i++ {
		minValue := (*ListaDesordenada)[i]
		minIndex := i
		for j := i; j < len(*ListaDesordenada); j++ {
			value := (*ListaDesordenada)[j]
			if value < minValue {
				minValue = value
				minIndex = j
			}
		}
		(*ListaDesordenada)[minIndex], (*ListaDesordenada)[i] = (*ListaDesordenada)[i], (*ListaDesordenada)[minIndex]
	}
	ch <- time.Since(start)
}

func main() {
	listRand := [][]int{rand.Perm(100), rand.Perm(1000), rand.Perm(1000000)}
	ch1 := make(chan time.Duration)
	ch2 := make(chan time.Duration)
	ch3 := make(chan time.Duration)

	for i := 0; i < 3; i++ {
		array := listRand[i]
		go Insertion(array, ch1)
		go Bubble(array, ch2)
		go Selection(&array, ch3)

		f.Println("Sorting time for :", len(array))
		f.Println("\tInsertion sort:", <-ch1)
		f.Println("\tBubble sort:", <-ch2)
		f.Println("\tSelection sort:", <-ch3)
	}

}
