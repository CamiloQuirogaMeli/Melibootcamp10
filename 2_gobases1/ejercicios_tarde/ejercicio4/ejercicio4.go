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
	switch mes{
        	case 1:
                	fmt.Printf("%d, Enero\n", mes)
                case 2:
                        fmt.Printf("%d, Febrero\n", mes)
		case 3:
                        fmt.Printf("%d, Marzo\n", mes)
		case 4:
                        fmt.Printf("%d, Abril\n", mes)
		case 5:
                        fmt.Printf("%d, Mayo\n", mes)
		case 6:
                        fmt.Printf("%d, Junio\n", mes)
		case 7:
                        fmt.Printf("%d, Julio\n", mes)
		case 8:
                        fmt.Printf("%d, Agosto\n", mes)
		case 9:
                        fmt.Printf("%d, Septiembre\n", mes)
		case 10:
                        fmt.Printf("%d, Octubre\n", mes)
		case 11:
                        fmt.Printf("%d, Noviembre\n", mes)
		case 12:
                        fmt.Printf("%d, Diciembre\n", mes)
	}
}
