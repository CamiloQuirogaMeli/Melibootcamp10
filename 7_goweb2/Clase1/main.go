package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var usuariosGlobales []Usuario
var err error
var lastId int

type Usuario struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Creacion time.Time `json:"creacion"`
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

func validacionCampos(u Usuario) (bool, string) {
	var camposInvalidos string
	completos := true

	sufijosValidosEmail := [3]string{".com", ".co", ".ar"}
	tieneSufijo := false

	if u.Nombre == "" {
		camposInvalidos += "nombre "
		completos = false
	}

	if u.Apellido == "" {
		camposInvalidos += ", apellido"
		completos = false
	}
	if u.Email == "" || strings.HasPrefix(u.Email, "@") || strings.Count(u.Email, "@") != 1 {
		camposInvalidos += "email"
		completos = false
	}

	for _, sufijo := range sufijosValidosEmail {
		if strings.HasSuffix(u.Email, sufijo) {
			tieneSufijo = true
			break
		}
	}

	if !tieneSufijo {
		camposInvalidos += ", email"
		completos = false
	}

	if u.Edad < 18 {
		camposInvalidos += ", edad"
		completos = false
	}
	if u.Altura < 0 {
		camposInvalidos += ", altura"
		completos = false
	}

	return completos, camposInvalidos

}

func PostUser(ctxt *gin.Context) {

	token := ctxt.Request.Header.Get("token")

	if token != "123456789" {
		ctxt.JSON(401, gin.H{
			"error": "No tiene permisos para realizar la petición solicitada",
		})

		return
	}

	var u Usuario
	if err := ctxt.Bind(&u); err != nil {
		ctxt.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastId++
	u.Id = lastId
	u.Activo = false
	u.Creacion = time.Now()

	validado, errores := validacionCampos(u)
	if validado {
		usuariosGlobales = append(usuariosGlobales, u)
	} else {
		log.Fatal("Los campos ", errores, " son erroneos")
	}

}

func main() {

	numUsuarios := 4
	lastId = numUsuarios
	usuariosGlobales, err = GenerarListaUsuarios(numUsuarios)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	users := router.Group("/users")

	users.GET("/", GetAll)

	users.POST("/crear", PostUser)

	router.Run()

}
