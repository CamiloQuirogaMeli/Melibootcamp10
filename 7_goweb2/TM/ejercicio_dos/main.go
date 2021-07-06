package main

import (
	"github.com/gin-gonic/gin"
)

type Silla struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Gamer  bool   `json:"gamer"`
	Color  string `json:"color"`
}

var sillasArr = make([]Silla, 0)

func getLastId() (value int) {
	value = 0
	for _, v := range sillasArr {
		if v.ID > value {
			value = v.ID
		}
	}
	value = value + 1
	return
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Silla

		if err := c.Bind(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Nombre == "" {
			c.JSON(400, gin.H{
				"error": "el campo nombre es obligatorio",
			})
			return
		}

		if req.Color == "" {
			c.JSON(400, gin.H{
				"error": "el campo color es obligatorio",
			})
			return
		}

		req.ID = getLastId()
		sillasArr = append(sillasArr, req)
		c.JSON(200, req)

	}
}

func main() {

	router := gin.Default()

	router.POST("/silla/", Guardar())

	router.Run()

}
