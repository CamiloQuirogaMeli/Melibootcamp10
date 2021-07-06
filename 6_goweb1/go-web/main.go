package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
)

type User struct {
	Id int
	Nombre string
	Apellido string
	Email string
	Edad int
	Altura float64
	Activo bool
	Creacion string
}

type Users struct {
	Usuarios []User
}

/* Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.

Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/

func main() {
	router := gin.Default()

	router.GET("/hola", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Gonzalo",
		})
	})

	router.GET("/usuarios", func (c *gin.Context) {
		sliceBytes, err := ioutil.ReadFile("./usuarios.json")

		/* Id int
		Nombre string
		Apellido string
		Email string
		Edad int
		Altura float64
		Activo bool
		Creacion string */
		
		id, nombre, apellido, email, edad, altura, activo, creacion := c.Query("id"), c.Query("nombre"), c.Query("apellido"),
		c.Query("email"), c.Query("edad"), c.Query("altura"), c.Query("activo"), c.Query("creacion")

		if err != nil {
			c.JSON(500, gin.H{
				"error": "No se pudo leer el archivo solicitado.",
			})
		} else {
			var users Users
			json.Unmarshal(sliceBytes, &users)
			if len(id) < 1 && len(nombre) < 1 && len(apellido) < 1 && len(email) < 1 && len(edad) < 1 && len(altura) < 1 && len(activo) < 1 && len(creacion) < 1 {
				c.String(200, "%v\nSIN FILTROS." , users)
			} else { // Aplicar filtros acá...
				var (
					filtered []User
					put bool
				)
				for i := 0; i <  len(users.Usuarios); i++ {
					put = true
					Id, _ := strconv.Atoi(id)
					if len(id) > 0 && users.Usuarios[i].Id != Id {
						put = false
						continue
					}
					if len(nombre) > 0 && strings.ToLower(users.Usuarios[i].Nombre) != strings.ToLower(nombre) {
						put = false
						continue
					}
					if len(apellido) > 0 && strings.ToLower(users.Usuarios[i].Apellido) != strings.ToLower(apellido) {
						put = false
						continue
					}
					if len(email) > 0 && strings.ToLower(users.Usuarios[i].Email) != strings.ToLower(email) {
						put = false
						continue
					}
					Edad, _ := strconv.Atoi(edad)
					if len(edad) > 0 && users.Usuarios[i].Edad != Edad {
						put = false
						continue
					}
					Altura, _ := strconv.ParseFloat(altura, 64)
					if len(altura) > 0 && users.Usuarios[i].Altura != Altura {
						put = false
						continue
					}
					Activo, _ := strconv.ParseBool(activo)
					if len(activo) > 0 && users.Usuarios[i].Activo != Activo {
						put = false
						continue
					}
					if len(creacion) > 0 && users.Usuarios[i].Creacion != creacion {
						put = false
						continue
					}

					if put == true { filtered = append(filtered, users.Usuarios[i]) }
				}
				c.String(200, "%v" , filtered)
			}
		}
	})

	router.GET("/usuarios/:id", func (c *gin.Context) {
		sliceBytes, err := ioutil.ReadFile("./usuarios.json")

		/* Id int
		Nombre string
		Apellido string
		Email string
		Edad int
		Altura float64
		Activo bool
		Creacion string */
		
		id := c.Param("id")

		if err != nil {
			c.JSON(500, gin.H{
				"error": "No se pudo leer el archivo solicitado.",
			})
		} else {
			var (
				users Users
				find bool
			)
			json.Unmarshal(sliceBytes, &users)
			
			for i := 0; i <  len(users.Usuarios); i++ {
				Id, _ := strconv.Atoi(id)
				if users.Usuarios[i].Id != Id {
					continue
				} else {
					find = true
					c.String(200, "%v" , users.Usuarios[i])
					break
				}
			}
			if !find {
				c.String(404, "No se encontro")
			}
		}
	})

	router.Run()
}