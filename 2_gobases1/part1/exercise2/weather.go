package main

import (
	"fmt"
)

func main() {
	
	//Decidi poner variables tipo int por que nos serian mas utiles los valores en este tipo, ya que si queremos obtener estadisticas
	//para analizar los cambios de clima es mas facil que un string.
	var temp int = 26
	var pressure int = 1018
	var wet int = 92

	fmt.Println("Temperature:" , temp, "°")
	fmt.Println("Pressure: ", pressure, "hPa")
	fmt.Println("Wet: ", wet, "%")
}
