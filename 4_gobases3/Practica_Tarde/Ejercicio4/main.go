package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Insercion(num []int, c chan time.Duration) {
	i := time.Now()
	var n = len(num)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if num[j-1] > num[j] {
				num[j-1], num[j] = num[j], num[j-1]
			}
			j = j - 1
		}
	}
	c <- time.Since(i)
}

func Burbuja(num []int, c chan time.Duration) {
	i := time.Now()
	n := len(num)
	cambio := true
	for cambio {
		cambio = false
		for i := 0; i < n-1; i++ {

			if num[i] > num[i+1] {

				num[i], num[i+1] = num[i+1], num[i]

				cambio = true
			}
		}
	}
	c <- time.Since(i)
}

func Seleccion(num []int, c chan time.Duration) {
	i := time.Now()
	var n = len(num)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if num[j] < num[minIdx] {
				minIdx = j
			}
		}
		num[i], num[minIdx] = num[minIdx], num[i]
	}
	c <- time.Since(i)
}

func main() {
	variable1 := make([]int, 100)
	variable1 = append(variable1, rand.Perm(100)...)

	variable2 := make([]int, 1000)
	variable2 = append(variable2, rand.Perm(1000)...)

	variable3 := make([]int, 10000)
	variable3 = append(variable3, rand.Perm(10000)...)

	c1 := make(chan time.Duration)
	c2 := make(chan time.Duration)
	c3 := make(chan time.Duration)

	go Insercion(variable1, c1)
	go Insercion(variable2, c2)
	go Insercion(variable3, c3)

	fmt.Println("Tiempo de Insercion con 100: ", <-c1)
	fmt.Println("Tiempo de Insercion con 1000: ", <-c2)
	fmt.Println("Tiempo de Insercion con 10000: ", <-c3)

	go Burbuja(variable1, c1)
	go Burbuja(variable2, c2)
	go Burbuja(variable3, c3)

	fmt.Println("Tiempo de Burbuja con 100: ", <-c1)
	fmt.Println("Tiempo de Burbuja con 1000: ", <-c2)
	fmt.Println("Tiempo de Burbuja con 10000: ", <-c3)

	go Seleccion(variable1, c1)
	go Seleccion(variable2, c2)
	go Seleccion(variable3, c3)

	fmt.Println("Tiempo de Seleccion con 100: ", <-c1)
	fmt.Println("Tiempo de Seleccion con 1000: ", <-c2)
	fmt.Println("Tiempo de Seleccion con 10000: ", <-c3)

}
