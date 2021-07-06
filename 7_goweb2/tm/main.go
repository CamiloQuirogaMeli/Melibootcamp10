package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID     int     `json:"id"`
	Name   string  `json:"nombre" binding:"required"`
	Type   string  `json:"tipo" binding:"required"`
	Amount int     `json:"cantidad" binding:"required"`
	Price  float32 `json:"precio" binding:"required"`
}

var products []request

func main() {
	router := gin.Default()
	pr := router.Group("/productos")
	pr.POST("", SaveProductHandler())
	router.Run()
}

func SaveProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		var req request
		if err := c.Bind(&req); err != nil {
			var missingFields []string
			if req.Amount == 0 {
				missingFields = append(missingFields, "cantidad")
			}
			if req.Name == "" {
				missingFields = append(missingFields, "nombre")
			}
			if req.Price == 0 {
				missingFields = append(missingFields, "precio")
			}
			if req.Type == "" {
				missingFields = append(missingFields, "tipo")
			}
			res := fmt.Sprintf("El/los campos faltantes son: %s", missingFields)
			c.JSON(400, gin.H{
				"error": res,
			})
			return
		}
		var lastId int
		if len(products) != 0 {
			lastId = products[len(products)-1].ID
		}
		req.ID = lastId + 1
		products = append(products, req)
		c.JSON(200, req)
	}

}
