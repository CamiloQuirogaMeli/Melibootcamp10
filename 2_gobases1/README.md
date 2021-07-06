# Ejercicio 1

```go
package main

import "fmt"

func main() {
	var name = "Jonathan"
	var address = "Posadas, Misiones, Argentina"

	fmt.Println("Nombre: ", name)
	fmt.Println("Dirección: ", address)
}
```

# Ejercicio 2

```go
package main

import "fmt"

func main() {
	var (
		temp     = 25.4
		humidity = 80.3
		pressure = 1001
	)

	fmt.Println("Temperatura: ", temp)
	fmt.Println("Humedad: ", humidity)
	fmt.Println("Presion: ", pressure)

}

```

# Ejercicio 3

## Código original:

```go 
var 1nombre string  // Incorrecto, los nombres de variables deben comenzar con letras.
var apellido string // Correcto
var int edad        // Correcto
1apellido := 6      // Incorrecto, los nombres de variables deben comenzar con letras, además, al declarar de esta forma estoy diciendole al compilador que 1apellido será un integer, inadecuado para guardar un apellido
var licencia_de_conducir = true // Incorrecto, se deben utilizar camellCase para nombrar variables
var estatura de la persona int  // Incorrecto, no se deben usar espacios en los nombres de variables
cantidadDeHijos := 2            // Correcto
```
## Código corregido:

```go 
var nombre string
var apellido string
var int edad
apellido := "Apellido"    
var licenciaDeConducir = true
var estaturaDeLaPersona float64
cantidadDeHijos := 2
```


# Ejercicio 4

## Código original:

```go
var apellido string = "Gomez"
var edad int = "35"
boolean := "false"
var sueldo string = 45857.90
var nombre string = "Julián";
```

## Código corregido:

```go
var apellido = "Gomez"
var edad = 35
boolean := false
var sueldo = 45857.90
var nombre = "Julián";
```