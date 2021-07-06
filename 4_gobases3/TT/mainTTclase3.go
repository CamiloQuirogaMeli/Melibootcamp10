package main

import "fmt"

func main() {

	fmt.Print("Ingrese ejercicio a resolver (1 o 2 o 3 o 4):\n")
	var ejercicioAresolver int
	fmt.Scanln(&ejercicioAresolver)

	switch ejercicioAresolver {
	case 1:
		fmt.Print("Ingrese nombre de usuario:\n")
		var usuarioName string
		fmt.Scanln(&usuarioName)
		fmt.Print("Ingrese apellido de usuario:\n")
		var usuarioLname string
		fmt.Scanln(&usuarioLname)
		fmt.Print("Ingrese edad de usuario:\n")
		var edad int
		fmt.Scanln(&edad)
		fmt.Println("Para iniciar sesión debe ingresar correo y luego contraseña")
		fmt.Print("Ingrese correo de usuario:\n")
		var correo string
		fmt.Scanln(&correo)
		fmt.Print("Ingrese contraseña:\n")
		var contrasena string
		fmt.Scanln(&contrasena)

		fmt.Print("Los datos ingresados fueron:\nNombre-> ", usuarioName, "\nApellido-> ", usuarioLname, "\nEdad-> ", edad, "\nCorreo-> ", correo, "\nContraseña-> ", contrasena, "\n")
		fmt.Print("Desea modificar algún dato?\t nombreyapellido o edad o correo o contrasena \n")
		var datoAModificar string
		fmt.Scanln(&datoAModificar)

		if datoAModificar == "no" {
			fmt.Println("Usuario ingresado correctamente")
		} else {
			var usuarioPuntero *Usuario
			u := Usuario{Nombre: usuarioName, Apellido: usuarioLname, Edad: edad, Correo: correo, Contrasena: contrasena}
			usuarioPuntero = &u
			procesoModificar(usuarioPuntero, datoAModificar)
		}
	case 2:
		u := UsuarioEcom{NombreUsEcom: "diego", ApellidoUEcom: "dileo", CorreoUEcom: "diegodileo3"}

		//var prodPointer *ProductosEcom = &p
		//var userPointer *UsuarioEcom = &u
		p := nuevoProducto("macBook", 450000)
		p1 := nuevoProducto("ipad", 40000)
		agregarProducto(&u, p, 2)
		agregarProducto(&u, p1, 4)

		fmt.Println("Desea borrar los productos del usuario? (si o no) ")
		var respuesta string
		fmt.Scanln(&respuesta)
		if respuesta == "si" {
			borrarProducto(&u)
		} else {
			fmt.Println("Programa finalizado")
		}
	case 3:
		//harcodeo de valores
		p := ProductsCP{NameP: "macBook", CostP: 45, AmountP: 1}
		p1 := ProductsCP{"ipad", 40, 3}
		p2 := ProductsCP{"iphone", 10, 5}
		todosP := []ProductsCP{p, p1, p2}

		s := ServicesCP{NameServ: "fix screen", CostServ: 40, Minutes: 20}
		s1 := ServicesCP{"fix batery", 20, 60}
		todosS := []ServicesCP{s, s1}

		m := ManteinmentCP{NameMant: "batery", CostMant: 35}
		m1 := ManteinmentCP{"screen", 10}
		todosM := []ManteinmentCP{m, m1}

		//canales para ejecutar funciones
		c := make(chan float64)
		go addProducts(todosP, c)
		go addServices(todosS, c)
		go addManteinance(todosM, c)

		var total float64
		total += <-c
		total += <-c
		total += <-c
		fmt.Println("El total de productos, servicios y mantenimiento es: ", total)
		fmt.Println("Fin programa")
	case 4:

	case 5:
		break
	}
}
