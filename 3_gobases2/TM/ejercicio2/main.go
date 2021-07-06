package main

import (
	"fmt"
)


func main()  {
	var num = []int{1,2,3, -8}

	fmt.Println(promPos(num))

}

func promPos( num []int ) int{
	var numFinal int = 0
	for _,numero := range num{
		if numero > 0{
			numFinal= numFinal + numero

		}else {
			fmt.Println("El alumno tiene una nota en negativo")
			break
		}
	}
    numFinal = numFinal/len(num)
	return numFinal
}