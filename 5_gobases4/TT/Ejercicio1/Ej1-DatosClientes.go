package main

import (

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Tipo string `json:"tipo"`
	Cantidad int `json:"cantidad"`
	Precio int `json:"precio"`
}

var products []producto //Array de la entidad en memoria que guarda las peticiones que se hagan

func Guardar() gin.HandlerFunc{

	return func(c *gin.Context){
		//Recibe token
		token := c.Request.Header.Get("token")
		if token != "123456"{
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return //Corta si token es invalido
		}
		//Recibe producto
		var prod producto
		if err := c.Bind(&prod); err != nil{
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return //Corta si hubo error
		}
		//Validar que todos los campos esten completos
		if prod.Nombre == ""{
			c.JSON(400, "el nombre del producto es requerido.")
		}else if prod.Tipo == ""{
			c.JSON(400, "el tipo del producto es requerido.")
		}else if prod.Cantidad == 0{
			c.JSON(400, "la cantidad del producto es requerida.")
		}else if prod.Precio == 0{
			c.JSON(400, "el precio del producto es requerido.")
		}else{
			prod.Id = crearID()

			c.JSON(200, prod)

			products = append(products, prod) //Sumo el prod a la variable global de memoria
		}
	}
}

func crearID() int{
	i := len(products)
	var proxId int
	if i == 0{
		proxId = 1
	}else{
		proxId = products[i-1].Id + 1
	}
	return proxId
}

func main(){
	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", Guardar())
	r.Run()


}
