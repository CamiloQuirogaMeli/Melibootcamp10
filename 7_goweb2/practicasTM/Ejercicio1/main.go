package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)
type Productos struct {
    Id             int `json:"id,omitempty"`
    Nombre         string `json:"nombre,omitempty"`
    Color          string `json:"color,omitempty"`
    Precio         float64 `json:"precio,omitempty"`
    Stock          int    `json:"stock,omitempty"`
    Codigo         string `json:"codigo,omitempty"`
    Publicado      bool   `json:"publicado,omitempty"`
    Fecha_creacion string `json:"fecha___creacion,omitempty"`
}

var producto []Productos

var maxId int =0
func main() {

    dat, err := ioutil.ReadFile("Productos.json")
    if err != nil {

        log.Println(err)
    } else {

        json.Unmarshal(dat, &producto)

    }

    router := gin.Default()
    router.GET("/", HandlerRaiz)

    // A handler for GET request on /example
    router.GET("/hello-camilo", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello Camilo!",
        }) // gin.H is a shortcut for map[string]interface{}
    })
    gopher :=router.Group("/Productos")
    {
        gopher.GET("/",Products)
        gopher.GET("/product/:id", Product)
        //gopher.GET("/:id",BuscarProduct)
    }
    router.POST("/productos", guardar)
    router.Run()
}
func guardar(c *gin.Context) {
    var req Productos
    if err := c.Bind(&req); err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }
    for _, product:= range producto {
       maxId=product.Id+1
    }
    req.Id = maxId
    producto = append(producto, req)
    c.JSON(200, req)
}

func HandlerRaiz (c *gin.Context){
    c.String(200, "¡Bienvenido al servidor interino bootcolombiano :3")
}
func Product (c *gin.Context){
    dat, err := ioutil.ReadFile("Productos.json")
    id := c.Param("id")
    if err != nil {
        c.JSON(500, gin.H{
            "error": "No se pudo leer el archivo",
        })
        log.Println(err)
    } else {

        json.Unmarshal(dat, &producto)
        //c.String(200, "%v", producto)
        var buscar bool
        for i :=0; i<len(producto); i++{
            Id, _ := strconv.Atoi(id)
            if producto[i].Id!= Id{
                continue
            }else{
                buscar = true
                c.String(200, "%v", producto[i])
            }
        }
        if !buscar{
            c.String(404,"No se encontro el producto")
        }
    }
}

func Products (c *gin.Context){
    dat, err := ioutil.ReadFile("Productos.json")

    id, nombre,color,precio,stock,codigo,publicado,fecha := c.Query("id"), c.Query("nombre"),c.Query("color"),
        c.Query("precio"), c.Query("stock"), c.Query("codigo"),c.Query("publicado"),c.Query("fecha_creacion")
    if err != nil {
        c.JSON(500, gin.H{
            "error": "No se pudo leer el archivo",
        })
        log.Println(err)
    } else {
        var filtrados []Productos
        var put bool


        json.Unmarshal(dat, &producto)
        if len(id)<1 && len(nombre)<1 && len(color)<1 && len(precio)<1 && len(stock)<1 && len(codigo)<1 && len(publicado)<1 && len(fecha)<1{
            c.String(200, "%v Sin filtros", producto)
        }else {//aplicar filtros
            //c.String(200, "%v", producto)

            for i := 0; i < len(producto); i++ {
                put = true
                Id, _ := strconv.Atoi(id)
                if len(id) > 0 && producto[i].Id != Id {
                    put = false
                    continue
                }
                if len(nombre) > 0 && strings.ToLower(producto[i].Nombre) != strings.ToLower(nombre) {
                    put = false
                    continue
                }
                if len(color) > 0 && strings.ToLower(producto[i].Color) != strings.ToLower(color) {
                    put = false
                    continue
                }
                if len(codigo) > 0 && strings.ToLower(producto[i].Codigo) != strings.ToLower(codigo) {
                    put = false
                    continue
                }
                Precio, _ := strconv.ParseFloat(precio, 64)
                if len(precio) > 0 && producto[i].Precio != Precio {
                    put = false
                    continue
                }
                Stock, _ := strconv.Atoi(stock)
                if len(stock) > 0 && producto[i].Stock != Stock {
                    put = false
                    continue
                }
                Publicado, _ := strconv.ParseBool(publicado)
                if len(publicado) > 0 && producto[i].Publicado != Publicado {
                    put = false
                    continue
                }
                if len(fecha) > 0 && producto[i].Fecha_creacion != fecha {
                    put = false
                    continue
                }
                if put == true {
                    filtrados = append(filtrados, producto[i])
                }
            }
            c.String(200,"%v", filtrados)
        }


    }
}


