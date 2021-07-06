package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenoInsercion(c chan int, array []int) {
	start := time.Now()
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	finish := time.Since(start)
	fmt.Printf("Insercion demoro %s \n", finish)
	c <- 1
}

func ordenoBubble(c chan int, array []int) {
	start := time.Now()
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	finish := time.Since(start)
	fmt.Printf("Bubble demoro %s \n", finish)
	c <- 1
}

func ordenoSeleccion(c chan int, array []int) {
	start := time.Now()
	for i := 0; i < len(array)-1; i++ {
		minimo := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minimo] {
				array[j], array[minimo] = array[minimo], array[j]
			}
		}
	}
	finish := time.Since(start)
	fmt.Printf("Selection demoro %s \n", finish)
	c <- 1
}

func main() {
	c := make(chan int)
	arr1 := rand.Perm(100)
	arr2 := rand.Perm(1000)
	arr3 := rand.Perm(1000000)
	fmt.Println("Ordenamientos para 100 numeros")
	go ordenoInsercion(c, arr1)
	go ordenoSeleccion(c, arr1)
	go ordenoBubble(c, arr1)
	for i := 0; i < 3; i++ {
		<-c
	}
	fmt.Println("Ordenamientos para 1000 numeros")
	go ordenoInsercion(c, arr2)
	go ordenoSeleccion(c, arr2)
	go ordenoBubble(c, arr2)
	for i := 0; i < 3; i++ {
		<-c
	}
	fmt.Println("Ordenamientos para 1000000 numeros")
	go ordenoInsercion(c, arr3)
	go ordenoSeleccion(c, arr3)
	go ordenoBubble(c, arr3)
	for i := 0; i < 3; i++ {
		<-c
	}

}
