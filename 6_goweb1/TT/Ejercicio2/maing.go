package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id        int
	Nombre    string
	Precio    int
	Publicado bool
	Color     string
	Fecha     string
	Stock     int
	Codigo    string
}

func GetAll(direccion string, p *[]product) []product {
	content, err := ioutil.ReadFile(direccion)
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(content, p)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Print(*p)
	return *p
}

func PaginaPrincipal(ctxt *gin.Context) {
	productos := []product{}
	ctxt.String(http.StatusOK, "%v", GetAll("../../TM/Ejercicio1/productos.json", &productos))
}

func BuscarProducto(ctxt *gin.Context) {
	productos := []product{}
	productos = GetAll("../../TM/Ejercicio1/productos.json", &productos)
	a := ctxt.Param("id")
	fmt.Print(a)
	//i, _ := strconv.Atoi(ctxt.Param("id"))
	producto := a
	var existe bool = false
	var productoFinal product

	for i := 0; i < len(productos) && !existe; i++ {
		ini := "id="
		if producto == fmt.Sprintf("%s%d", ini, productos[i].Id) {
			existe = true
			productoFinal = productos[i]
		}

	}

	if existe {
		ctxt.String(200, "Información del producto %v, nombre: %s", ctxt.Param("id"), productoFinal)
	} else {
		//De no existir, responde con 404 - NOT FOUND
		ctxt.String(404, "Información del producto ¡No existe!")
	}
}

func main() {

	server := gin.Default()
	server.GET("/productos", PaginaPrincipal)
	server.GET("/productos/:id", BuscarProducto)
	server.Run(":8080")
}
