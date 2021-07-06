package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}
	for i := 0; i < 10; i++ {

		fmt.Println("se ha terminado el programa en :", <-c)

	}
}

func proceso(i int, c chan int) {
	fmt.Println(i, "-inicio")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Terminado")
	c <- i

}
