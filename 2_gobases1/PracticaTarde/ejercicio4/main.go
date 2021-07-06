package main

import "fmt"

func main() {
	var nroMes uint8
	var meses = map[uint8]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Ocutbre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Println("Ingrese el nro de mes que desea consultar:")
	fmt.Scanln(&nroMes)
	fmt.Println(meses[nroMes])
	/*tambien se puede resolver con un switch donde evaluo el nro ingresado y segun cual sea el nro
	  ingresa al case donde hay un println con el mes textual
	  ejemplo: case 2 fmt.Println("Febrero")*/
}
