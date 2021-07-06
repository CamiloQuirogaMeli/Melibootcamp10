package main

import (
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float32 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     string  `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var ListaProductos = []Producto{
	{
		Id:            1,
		Nombre:        "Producto 1",
		Color:         "Rojo",
		Precio:        753.47,
		Stock:         54,
		Codigo:        "1-P1",
		Publicado:     "Si",
		FechaCreacion: "20/03/2021",
	},
	{
		Id:            2,
		Nombre:        "Producto 2",
		Color:         "Negro",
		Precio:        943.47,
		Stock:         21,
		Codigo:        "2-P2",
		Publicado:     "Si",
		FechaCreacion: "25/03/2021",
	},
	{
		Id:            3,
		Nombre:        "Producto 3",
		Color:         "Azul",
		Precio:        900.47,
		Stock:         13,
		Codigo:        "3-P3",
		Publicado:     "Si",
		FechaCreacion: "10/03/2021",
	},
	{
		Id:            4,
		Nombre:        "Producto 4",
		Color:         "Amarillo",
		Precio:        500.47,
		Stock:         5,
		Codigo:        "4-P4",
		Publicado:     "Si",
		FechaCreacion: "20/03/2021",
	},
	{
		Id:            5,
		Nombre:        "Producto 5",
		Color:         "Violeta",
		Precio:        912.47,
		Stock:         54,
		Codigo:        "5-P5",
		Publicado:     "Si",
		FechaCreacion: "20/03/2021",
	},
}

func main() {
	server := gin.Default()
	productos := server.Group("/productos")
	{
		productos.GET("/GetAll", handlerGetAll)
		productos.POST("/", handlerPost)
	}

	server.Run()
}

// GO Web 1 TM - Ej1
func handlerGetAll(ctx *gin.Context) {
	ctx.JSON(200, ListaProductos)
}

// Go Web 2 TM - Ej 1 - Ej3
func handlerPost(ctx *gin.Context) {
	//Go Web 2 TM - Ej3
	if ctx.Request.Header.Get("Token") != "productEndpoint" {
		ctx.JSON(401, "No tiene permisos para realizar la petición solicitada")
		return
	}

	var prod Producto
	if err := ctx.Bind(&prod); err != nil {
		ctx.JSON(404, gin.H{
			"error": err,
		})
		return
	}
	mensajeError := validarProducto(prod)
	if mensajeError != "" {
		ctx.JSON(400, mensajeError)
		return
	}
	prod.Id = generarId(prod)
	ListaProductos = append(ListaProductos, prod)
	ctx.JSON(200, prod)
}

// Ej 1 GoWeb2 TM
func generarId(prod Producto) int {
	if len(ListaProductos) > 0 {
		return ListaProductos[len(ListaProductos)-1].Id + 1
	} else {
		return 1
	}
}

//Go Web 2 - Ej 2
func validarProducto(prod Producto) string {
	switch {
	case prod.Nombre == "":
		return "El nombre es requerido"
	case prod.Color == "":
		return "El color es requerido"
	case prod.Precio == 0:
		return "El precio es requerido"
	case prod.Stock == 0:
		return "El stock es requerido"
	case prod.Codigo == "":
		return "El codigo es requerido"
	case prod.Publicado == "":
		return "El campo punblicado es requerido"
	case prod.FechaCreacion == "":
		return "La fecha de creacion es requerida"
	}
	return ""
}
