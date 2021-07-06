/*
Realizar una aplicación que contenga una variable con el número del mes, imprimir el mes
que corresponda. ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y
por qué?
Ej: 7, Julio
*/

// package declaration
package main

// library imports
import (
	"fmt"
	"os"
	"time"
)

func main() {
	options := `	1. Ver la fecha completa
	2. Ver el mes actual
	3. Ver el dia actual
	4. Ver el año actual
	5. Ver la fecha en formato (año mes dia)
	6. Convertir el numero de mes en el nombre del mismo
	7. Convertir el numero del dia de la semana en el nombre del mismo
	8. Ejecutar un script despues de 10 segundos
	9. Salir`
	continous := 0
	for continous == 0 {
		fmt.Println("Elija la opcion que desee realizar")
		fmt.Println(options)
		var selectItem int
		fmt.Print("-> ")
		_, errProduct := fmt.Scanf("%d", &selectItem)

		if errProduct != nil {
			fmt.Println("Bad Input")
			os.Exit(0)
		}
		fmt.Println()
		switch selectItem {
		case 1:
			fmt.Println(time.Now())
		case 2:
			fmt.Println(time.Now().Month())
		case 3:
			fmt.Println(time.Now().Day())
		case 4:
			fmt.Println(time.Now().Year())
		case 5:
			fmt.Println(time.Now().Date())
		case 6:
			var numberMonth int
			fmt.Print("Por favor digite el numero-> ")
			_, errNumber := fmt.Scanf("%d", &numberMonth)

			if errNumber != nil || numberMonth > 12 {
				fmt.Println("Bad Input")
				os.Exit(0)
			}
			fmt.Println(time.Month(numberMonth))
		case 7:
			var numberDay int
			fmt.Print("Por favor digite el numero-> ")
			_, errNumber := fmt.Scanf("%d", &numberDay)

			if errNumber != nil || numberDay > 6 {
				fmt.Println("Bad Input")
				os.Exit(0)
			}
			fmt.Println(time.Weekday(numberDay))
		case 8:

			time.AfterFunc(10*time.Second, func() {
				fmt.Println("Script ejecutado despues de 10 segundos")
			})
			time.Sleep(10 * time.Second)
		case 9:
			os.Exit(0)
		default:
			os.Exit(0)
		}
		fmt.Println()
		fmt.Println("-------------------------------------")
		fmt.Println()
	}
}
