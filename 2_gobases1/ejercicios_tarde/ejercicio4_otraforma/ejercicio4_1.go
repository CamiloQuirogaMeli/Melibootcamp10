package main

import(
	"fmt"
)

func main(){
	var mes int16 = 0
        for mes < 1 || mes > 12{
                        fmt.Print("Ingrese un número de mes válido (del 1 al 12): ")
                        fmt.Scanln(&mes)
                }
	var meses = map[int16]string{
		1:"Enero",
		2:"Febrero",
		3:"Marzo",
		4:"Abril",
		5:"Mayo",
		6:"Junio",
		7:"Julio",
		8:"Agosto",
		9:"Septiembre",
		10:"Octubre",
		11:"Noviembre",
		12:"Diciembre"}
	for key, element := range meses {
 	if key == mes{
		fmt.Println(key, ",", element)
		}
	}
}
