package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	NOMBRE   = "Julián"
	APELLIDO = "Velandia"
)

var usuariosGlobales []Usuario
var err error

type Usuario struct {
	Id       int
	Nombre   string
	Apellido string
	Email    string
	Edad     int
	Altura   float64
	Activo   bool
	Creacion time.Time
}

func (u *Usuario) inicializarAleatorio(id int) {
	nombres := [4]string{"Juan", "Laura", "Esteban", "Sofia"}

	u.Id = id
	u.Nombre = nombres[rand.Intn(4)]
	u.Apellido = nombres[rand.Intn(4)]
	u.Email = u.Nombre + "@correo.com"
	u.Edad = rand.Intn(40) + 18
	u.Altura = rand.Float64()*10 + 150
	u.Activo = rand.Intn(2) != 0
	u.Creacion = time.Now()

}

func GenerarListaUsuarios(numEmpleados int) ([]Usuario, error) {
	listaUsuarios := make([]Usuario, 0)

	if numEmpleados < 0 {
		return listaUsuarios, errors.New("número negativo")
	}
	for i := 0; i < numEmpleados; i++ {
		u := Usuario{}
		u.inicializarAleatorio(i)
		listaUsuarios = append(listaUsuarios, u)
	}

	return listaUsuarios, nil
}

func GetAll(c *gin.Context) {

	var datos string

	for _, usuario := range usuariosGlobales {
		datosJson, err := json.Marshal(usuario)
		if err != nil {
			log.Fatal(err)
		}
		datos += string(datosJson)

	}

	c.String(200, datos)
}

func Enviarsaludo(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "Hola " + NOMBRE,
	})
}

func GetFilteredNombre(ctxt *gin.Context) {
	var filtrados []Usuario

	for _, e := range usuariosGlobales {
		if ctxt.Param("Nombre") == e.Nombre {
			filtrados = append(filtrados, e)
		}
	}

	GetFiltered(filtrados, ctxt)
}

func GetFilteredActivo(ctxt *gin.Context) {

	var filtrados []Usuario

	for _, e := range usuariosGlobales {
		if ctxt.Param("Activo") == strconv.FormatBool(e.Activo) {
			filtrados = append(filtrados, e)
		}
	}

	GetFiltered(filtrados, ctxt)
}

func GetFilteredEdad(ctxt *gin.Context) {

	var filtrados []Usuario

	for _, e := range usuariosGlobales {
		if ctxt.Param("Edad") == strconv.Itoa(e.Edad) {
			filtrados = append(filtrados, e)
		}
	}

	GetFiltered(filtrados, ctxt)
}

func GetFilteredId(ctxt *gin.Context) {

	var filtrados []Usuario

	for _, e := range usuariosGlobales {
		if ctxt.Param("Id") == strconv.Itoa(e.Id) {
			filtrados = append(filtrados, e)
		}
	}

	GetFiltered(filtrados, ctxt)
}

func GetFiltered(filtrados []Usuario, ctxt *gin.Context) {

	datos := "Datos \n"
	if len(filtrados) > 0 {
		for _, usuario := range filtrados {
			datosJson, err := json.Marshal(usuario)
			if err != nil {
				log.Fatal(err)
			}
			datos += string(datosJson)

		}
		ctxt.String(200, datos)
	} else {
		ctxt.String(404, "No hay usuarios con esos datos")
	}

}

func main() {

	numUsuarios := 4

	usuariosGlobales, err = GenerarListaUsuarios(numUsuarios)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/hola", Enviarsaludo)

	users := router.Group("/users")
	{
		users.GET("/", GetAll)
		users.GET("/:Id", GetFilteredId)
		users.GET("/filtered_nombre/:Nombre", GetFilteredNombre)
		users.GET("/filtered_edad/:Edad", GetFilteredEdad)
		users.GET("/filtered_activo/:Activo", GetFilteredActivo)

	}

	router.Run()

}
