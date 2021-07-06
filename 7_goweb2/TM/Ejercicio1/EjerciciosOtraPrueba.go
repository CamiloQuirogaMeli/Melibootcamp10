package main

import(
	"github.com/gin-gonic/gin"
)
//En esta prueba agrego el `binding:"required"` dentro del struct
//para aquellos parametros obligatorios (funcionalidad dentro de paquete de gin)

type producto struct {
	Id int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Tipo string `json:"tipo" binding:"required"`
	Cantidad int `json:"cantidad" binding:"required"`
	Precio int `json:"precio" binding:"required"`
}

var productos []producto

func Guardar() gin.HandlerFunc{
	return func(c *gin.Context){
		//Token
		token := c.Request.Header.Get("token")
		if token != "123456"{
			c.JSON(401, gin.H{
				"error": "token invalido.",
			})
			return //En caso de error corta
		}
		//Producto
		var prod producto
		if err := c.Bind(&prod); err != nil{
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return //Corta en caso de error
		}
		prod.Id = crearID()
		c.JSON(200, prod)
		productos = append(productos, prod)

	}
}

func crearID() int {
	var proxID int
	i := len(productos)
	if len(productos) == 0{
		proxID = 1
	}else{
		proxID = productos[i-1].Id + 1
	}
	return proxID
}

func main(){
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	r.Run()
}