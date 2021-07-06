package main

import (
	"github.com/gin-gonic/gin"
	"errors"
	/*"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"*/
)

/* Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global). */

/* Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo). */

/* Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”. */

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

type request struct {
	id int `json:"id"`
	nombre string `json:"nombre"`
	apellido string `json:"apellido"`
	email string `json:"email"`
	edad int `json:"edad"`
	altura float64 `json:"altura"`
	activo bool `json:"activo"`
	creacion string `json:"creacion"`
}

var Usuarios []User

func joinUser(id int, nombre, apellido, email string, edad int, altura float64, activo bool, creacion string) error {
	if len(nombre) < 1 || len(apellido) < 1 || len(email) < 1 || edad < 1 || altura < 0.5 || len(creacion) < 1 {
		return errors.New("la creacion de un usuario debe tener todos los campos.")
	}
	Usuarios = append(Usuarios, User{id, nombre, apellido, email, edad, altura, activo, creacion})
	return nil
}

func main() {
	router := gin.Default()

	router.POST("/user", func (c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			req.id = len(Usuarios)+1
			err := joinUser(req.id, req.nombre, req.apellido, req.email, req.edad, req.altura, req.activo, req.creacion)
			if err != nil {
				c.JSON(200, gin.H{
					"error":err.Error(),
				})
			} else {
				c.JSON(202, req)
			}
		}
	})

	router.Run()
}