package main

import(
        "fmt"
	"strings"
        )

func main(){
        fmt.Print("Ingrese su edad: ")
        var e int32
        fmt.Scanln(&e)
        if e < 22{
                fmt.Println("Para acceder a nuestros prestamos debe ser mayor de 22 años")
        }else{
		var empleado string = ""
		for strings.ToUpper(empleado) != "S" && strings.ToUpper(empleado) != "N"{
			fmt.Print("Es usted empleado en relacion de dependencia? (S = Si, N = No)")
        	   	fmt.Scanln(&empleado)
		}
		if strings.ToUpper(empleado) == "N"{
			fmt.Println("Para acceder a nuestros prestamos debe ser empleado")
		}else{
			var antiguedad float32
                	fmt.Print("Indique los años de antiguedad en la empresa: ")
                        fmt.Scanln(&antiguedad)
			if(antiguedad < 1){
				fmt.Println("Para acceder a nuestros prestamos debe tener, al menos, 1 año de antigüedad")
			}else{
				var sueldo float32
                        	fmt.Print("Ingrese su sueldo:")
                        	fmt.Scanln(&sueldo)
				switch{
				case sueldo >= 100000:
					fmt.Println("Usted tiene a disposición un crédito")
				case sueldo > 0 && sueldo < 100000:
					fmt.Println("Usted tiene a disposición un crédito a tasa 0%")
				default:
					fmt.Println("El sueldo no puede ser negativo")
				}
			}

		}
	}
}
