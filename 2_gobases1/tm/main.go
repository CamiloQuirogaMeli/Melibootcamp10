package main

import (
	"fmt"
	"meli_bootcamp10/2_gobases1/tm/ejercicios"
)

func main() {
	fmt.Println("Primer ejercicio")
	ejercicios.PrimerEjercicio()

	fmt.Println("\nSegundo ejercicio")
	ejercicios.SegundoEjercicio()

	fmt.Println("\nTercer ejercicio")
	ejercicios.TercerEjercicio()

	fmt.Println("\nCuarto ejercicio")
	const MES = 6
	ejercicios.CuartoEjercicio(MES)
	ejercicios.CuartoEjercicioVersionDos(MES)
	ejercicios.CuartoEjercicioVersionTres(MES)

	fmt.Println("\nQuinto ejercicio - literal A")
	ejercicios.QuintoEjercicioLiteralA()
	fmt.Println("\nQuinto ejercicio - literal B")
	ejercicios.QuintoEjercicioLiteralB()

	fmt.Println("\nSexto ejercicio")
	ejercicios.SextoEjercicio()
}
