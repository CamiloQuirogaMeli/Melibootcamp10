package main

import "fmt"

type Usuario struct {
	nombre            string
	edad              int
	empleado          bool
	antiguedadEnMeses int
	sueldo            float64
}

func main() {

	usuario1 := nuevoUsuario("Carlos", 21, true, 14, 85000)
	usuario2 := nuevoUsuario("Maria", 45, true, 50, 105000)
	usuario3 := nuevoUsuario("Marcelo", 23, false, 15, 60000)
	usuario4 := nuevoUsuario("Marta", 31, true, 9, 50000)
	usuario5 := nuevoUsuario("Florencio", 31, true, 65, 90000)

	var usuarios = []Usuario{usuario1, usuario2, usuario3, usuario4}

	usuarios = append(usuarios, usuario5)

	for i, usuario := range usuarios {
		fmt.Println(i + 1)
		fmt.Println("Usuario: ", usuario.nombre)
		fmt.Println(validarUsuario(usuario))
	}
}

func nuevoUsuario(nombre string, edad int, empleado bool, antiguedadEnMeses int, sueldo float64) Usuario {
	u := new(Usuario)
	u.nombre = nombre
	u.edad = edad
	u.empleado = empleado
	u.antiguedadEnMeses = antiguedadEnMeses
	u.sueldo = sueldo
	return *u
}

func validarUsuario(usuario Usuario) string {
	switch {
	case usuario.edad < 23:
		return "No tiene la edad suficiente para recibir un préstamo"
	case !usuario.empleado:
		return "No se otorgan préstamos a usuarios desempleados"
	default:
		if usuario.sueldo > 100000 {
			return "Préstamo CON INTERÉS otorgado. ¡Felicitaciones!"
		} else {
			return "Préstamo SIN INTERÉS otorgado. ¡Felicitaciones!"
		}
	}
}
