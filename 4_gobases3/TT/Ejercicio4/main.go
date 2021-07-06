package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int, c chan int) []int {
	t0 := time.Now()
	var flag bool
	for i := 0; i < len(arr)-1; i++ {
		flag = false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	t1 := time.Now()
	c <- int(t1.Sub(t0))
	return arr
}

func selectSort(arr []int, c chan int) []int {
	t0 := time.Now()
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	t1 := time.Now()
	c <- int(t1.Sub(t0))
	return arr
}

func insertSort(arr []int, c chan int) []int {
	t0 := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	t1 := time.Now()
	c <- int(t1.Sub(t0))
	return arr
}

func main() {
	c := make(chan int)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(1000000)

	go bubbleSort(variable1, c)
	go selectSort(variable1, c)
	go insertSort(variable1, c)

	fmt.Println("Arrays de 100 números")
	fmt.Println("El tiempo que demoró el ordenamiento por bubujas fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por selección fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por inserción fue de: ", <-c, "milisegundos")
	fmt.Println("----------------------------")

	go bubbleSort(variable2, c)
	go selectSort(variable2, c)
	go insertSort(variable2, c)

	fmt.Println("Arrays de 1000 números")
	fmt.Println("El tiempo que demoró el ordenamiento por bubujas fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por selección fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por inserción fue de: ", <-c, "milisegundos")
	fmt.Println("----------------------------")

	go bubbleSort(variable3, c)
	go selectSort(variable3, c)
	go insertSort(variable3, c)

	fmt.Println("Arrays de 100000 números")
	fmt.Println("El tiempo que demoró el ordenamiento por bubujas fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por selección fue de: ", <-c, "milisegundos")
	fmt.Println("El tiempo que demoró el ordenamiento por inserción fue de: ", <-c, "milisegundos")
	fmt.Println("----------------------------")

}
