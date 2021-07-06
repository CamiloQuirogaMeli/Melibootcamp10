package main

import "fmt"

func main() {
	// fmt.Println("Su salario es: ", calcTax())
	// fmt.Println("Su calificacion promedio es de: ", avgStudent(9, 9, 10))
	// fmt.Println("Su salario segun su categoria es de: ", calcSalary(4000, "A"))

	// studentCalculo, error := studentCalc()
	// fmt.Println("Calculo por estudiante", studentCalculo(9, 9, 10, 8, 7), error)

	// var animal string
	// fmt.Println("Indique que animal tiene(perro, gato, hamster, tarantula): ")
	// fmt.Scan(&animal)
	// foodPet, _ := calcFoodPet(animal)
	// fmt.Println(foodPet())

	// Ejercicios Tarde
	// alumno1 := Alumno{"Juan", "Perez", "12345678", "19/02/2021"}
	// alumno1.detalle()

	// matrix1 := Matrix{}
	// matriz := []float64{2.2, 2.4, 2.5, 10.5}
	// matrix1.set(matriz)
	// matrix1.print()

	// productoC := factory("Chico", 100)
	// fmt.Println("Producto chico: ", productoC.precio())
	// productoM := factory("Mediano", 100)
	// fmt.Println("Producto mediano: ", productoM.precio())
	// productoG := factory("Grande", 100)
	// fmt.Println("Producto grande: ", productoG.precio())

	prodChic := Chico4{Producto4{1000.0}}
	prodMed := Mediano4{Producto4{10000.0}}
	prodGra := Grande4{Producto4{10000000.0}}
	prodFragil := Fragil4{Producto4{10000000.0}}
	prodEspecial1 := Especial4{Producto4{1000000.0}}
	prodEspecial2 := Especial4{Producto4{1000000.0}}

	flete1 := Flete{[]IProducto{prodChic, prodMed, prodGra, prodFragil, prodEspecial1}}
	fmt.Println("Productos en el flete: ", flete1.productos)
	flete1.AgregarProducto(prodEspecial2)
	fmt.Println("Productos en el flete: ", flete1.productos)
	env, envEsp := flete1.CalcularEnvios()
	fmt.Println("Envios normales a realizar: ", env, ". Envios especiales a realizar: ", envEsp)
}
