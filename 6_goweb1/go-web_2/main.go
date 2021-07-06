package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ARCHIVO = "../go-web/product.json"

type Producto struct {
	Id             int
	Nombre         string
	Color          string
	Precio         float64
	Stock          int
	Codigo         string
	Publicado      bool
	Fecha_creacion string
}

func getAll(ctx *gin.Context) {
	data, err := ioutil.ReadFile(ARCHIVO)
	if err != nil {
		log.Fatal(err)
	}

	var p []Producto
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"Productos": p,
	})
}

func addProduc() ([]Producto, error) {
	data, err := ioutil.ReadFile(ARCHIVO)
	if err != nil {
		return []Producto{}, errors.New("error al leer el archivo")
	}
	var p []Producto
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		return []Producto{}, errors.New("error al cargar el archivo")
	}
	return p, nil
}

func getOne(ctx *gin.Context) {
	prod, err := addProduc()
	if err != nil {
		log.Fatal(err)
	}
	paramId := ctx.Param("Id")
	for _, values := range prod {
		if strconv.Itoa(values.Id) == paramId {
			ctx.JSON(200, values)
			return
		}
	}
	ctx.String(404, "Not found")
}

func getFilter(ctx *gin.Context) {
	prod, err := addProduc()
	if err != nil {
		log.Fatal(err)
	}
	var selectProd []Producto
	for _, values := range prod {
		if ctx.Query("Id") == strconv.Itoa(values.Id) || ctx.Query("Nombre") == values.Nombre || ctx.Query("Color") == values.Color ||
			ctx.Query("Precio") == strconv.Itoa(int(values.Precio)) || ctx.Query("Stock") == strconv.Itoa(values.Stock) || ctx.Query("Codigo") == values.Codigo ||
			ctx.Query("Publicado") == strconv.FormatBool(values.Publicado) || ctx.Query("Fecha_creacion") == values.Fecha_creacion {

			selectProd = append(selectProd, values)
		}
	}
	ctx.JSON(200, selectProd)
}

func main() {
	router := gin.Default()

	group := router.Group("/product")
	{
		group.GET("/", getAll)
		group.GET("/:Id", getOne)
		group.GET("/filterProduct", getFilter)
	}

	router.Run()
}
