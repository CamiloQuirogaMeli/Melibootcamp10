package main
import "fmt"

func main(){
        var temperatura int16 = 12
        var humedad uint16 = 81
        var presion float32 = 1015.6
        fmt.Println("La temperatura en Santa Fe es: ", temperatura, "ºC")
        fmt.Println("La humedad es: ", humedad, "%")
        fmt.Println("La presion atmosférica es:", presion, "mb")
}
