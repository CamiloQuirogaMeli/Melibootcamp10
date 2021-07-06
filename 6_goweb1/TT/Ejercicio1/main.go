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
	params := ctxt.Params
	fmt.Print(params)
	productos := []product{}
	ctxt.String(http.StatusOK, "%v", GetAll("../../TM/Ejercicio1/productos.json", &productos))
}

func main() {

	server := gin.Default()
	server.GET("/productos", PaginaPrincipal)
	server.GET("/productos/:publicado/:codigo", FiltrarProductosPublicados)
	server.Run(":8080")
}

func FiltrarProductosPublicados(ctxt *gin.Context) {
	productos := []product{}
	productos = GetAll("../../TM/Ejercicio1/productos.json", &productos)
	var filtrados []*product
	cantidadFiltrados := 0
	a := ctxt.Param("publicado")
	fmt.Print(a)
	for _, p := range productos {
		var estaPublicado string = ctxt.Query("publicado")
		if estaPublicado == "true" {
			filtrados = append(filtrados, &p)
			cantidadFiltrados++
		}
	}

	if cantidadFiltrados > 0 {
		ctxt.String(200, "¡Filtrado exitoso!\n")
		for _, p := range filtrados {
			log.Println("====== Filter Por Query String ======")
			log.Println(p.Id)
			log.Println(p.Nombre)
			log.Println(p.Publicado)
		}
	} else {
		ctxt.String(404, "¡Oh no! No hay productos no publicados\n")
		ctxt.String(http.StatusOK, "%v", productos)
	}

}

// func FiltrarProductosId(ctxt *gin.Context) {
// 	productos := []product{}
// 	productos = GetAll("../../TM/Ejercicio1/productos.json", &productos)
// 	var filtrados []*product
// 	cantidadFiltrados := 0
// 	for _, p := range productos {
// 		productoId := ctxt.Query("id")
// 		if productoId == string(p.Id) {
// 			filtrados = append(filtrados, &p)
// 			cantidadFiltrados++
// 		}
// 	}

// 	if cantidadFiltrados > 0 {
// 		ctxt.String(200, "¡Filtrado exitoso!\n")
// 		for _, p := range filtrados {
// 			log.Println("====== Filter Por Query String ======")
// 			log.Println(p.Id)
// 			log.Println(p.Nombre)
// 			log.Println(p.Publicado)
// 			log.Println(p.Color)
// 		}
// 	} else {
// 		ctxt.String(404, "¡Oh no! No hay productos con ese Id\n")
// 		ctxt.String(http.StatusOK, "%v", productos)
// 	}

// }

// func BuscarProducto(ctxt *gin.Context) {
// 	productos := []product{}
// 	productos = GetAll("../../TM/Ejercicio1/productos.json", &productos)
// 	a := ctxt.Param("id")
// 	fmt.Print(a)
// 	//i, _ := strconv.Atoi(ctxt.Param("id"))
// 	producto := a
// 	var existe bool = false
// 	var productoFinal product

// 	for i := 0; i < len(productos) && !existe; i++ {
// 		ini := "id="
// 		if producto == fmt.Sprintf("%s%d", ini, productos[i].Id) {
// 			existe = true
// 			productoFinal = productos[i]
// 		}

// 	}

// 	if existe {
// 		ctxt.String(200, "Información del producto %v, nombre: %s", ctxt.Param("id"), productoFinal)
// 	} else {
// 		//De no existir, responde con 404 - NOT FOUND
// 		ctxt.String(404, "Información del producto ¡No existe!")
// 	}
// }

// func FiltrarProductosPrecio(ctxt *gin.Context) {
// 	productos := []product{}
// 	productos = GetAll("../../TM/Ejercicio1/productos.json", &productos)
// 	var filtrados []*product
// 	cantidadFiltrados := 0
// 	for _, p := range productos {
// 		i, _ := strconv.Atoi(ctxt.Query("precio"))
// 		if i <= p.Precio {
// 			filtrados = append(filtrados, &p)
// 			cantidadFiltrados++
// 		}
// 	}

// 	if cantidadFiltrados > 0 {
// 		ctxt.String(200, "¡Filtrado exitoso!\n")
// 		for _, p := range filtrados {
// 			log.Println("====== Filter Por Query String ======")
// 			log.Println(p.Id)
// 			log.Println(p.Nombre)
// 			log.Println(p.Publicado)
// 			log.Println(p.Color)
// 		}
// 	} else {
// 		ctxt.String(404, "¡Oh no! No hay productos más baratos\n")
// 		ctxt.String(http.StatusOK, "%v", productos)
// 	}

// }
