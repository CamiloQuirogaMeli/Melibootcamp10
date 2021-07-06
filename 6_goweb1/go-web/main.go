package main

import (
	"strconv"

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
	// Ej 2 GoWeb TM
	server.GET("/hello-jhon", func(c *gin.Context) {
		c.JSON(200,
			"Hola Juan!",
		)
	})

	productos := server.Group("/productos")
	{
		productos.GET("/GetAll", handlerGetAll)
		productos.GET("/:Id", handlerFilter)
		productos.GET("/Get", handlerProducto)

	}

	server.Run()
}

// Ej 3 GoWeb TM
func handlerGetAll(ctx *gin.Context) {
	ctx.JSON(200, ListaProductos)
}

// GO Web 1 TT - Ej 2
func handlerFilter(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("Id"))
	if err != nil || id < 0 {
		ctx.JSON(400, "No se pudo obtener los productos: Id inválido")
		return
	}
	for key, producto := range ListaProductos {
		if producto.Id == int(id) {
			ctx.JSON(200, key)
			return
		}
	}
}

// GO Web 1 TT - Ej 1
func handlerProducto(ctx *gin.Context) {
	var resultado []Producto
	var query Producto
	if ctx.Bind(&query) == nil {
		switch {
		case query.Id != 0:
			for _, producto := range ListaProductos {
				if producto.Id == query.Id {
					resultado = append(resultado, producto)
				}
			}
		case query.Nombre != "":
			for _, producto := range ListaProductos {
				if producto.Nombre == query.Nombre {
					resultado = append(resultado, producto)
				}
			}
		case query.Color != "":
			for _, producto := range ListaProductos {
				if producto.Color == query.Color {
					resultado = append(resultado, producto)
				}
			}
		case query.Precio != 0.0:
			for _, producto := range ListaProductos {
				if producto.Color == query.Color {
					resultado = append(resultado, producto)
				}
			}
		case query.Stock != 0:
			for _, producto := range ListaProductos {
				if producto.Stock == query.Stock {
					resultado = append(resultado, producto)
				}
			}
		case query.Codigo != "":
			for _, producto := range ListaProductos {
				if producto.Codigo == query.Codigo {
					resultado = append(resultado, producto)
				}
			}
		case query.Publicado != "":
			for _, producto := range ListaProductos {
				if producto.Publicado == query.Publicado {
					resultado = append(resultado, producto)
				}
			}
		case query.FechaCreacion != "":
			for _, producto := range ListaProductos {
				if producto.FechaCreacion == query.FechaCreacion {
					resultado = append(resultado, producto)
				}
			}
		default:
			ctx.JSON(404, "No hay filtro")
		}
		if len(resultado) > 0 {
			ctx.JSON(200, resultado)
		} else {
			ctx.JSON(404, "Productos no encontrados")
		}

		return
	}
	ctx.JSON(404, "Llamada inválida")
}
