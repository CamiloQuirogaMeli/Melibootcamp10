package handler

import (
	"github.com/conavarr/tt/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type ProductController struct {
	service products.Service
}

func NewProductController(s products.Service) *ProductController {
	return &ProductController{
		service: s,
	}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if validateToken(ctx, token) {
			return
		}

		prods, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, prods)
	}
}

func validateToken(ctx *gin.Context, token string) bool {
	if token != "1234456" {
		ctx.JSON(401, gin.H{
			"error": "Token inválido",
		})
		return true
	}
	return false
}

func (c *ProductController) Store() gin.HandlerFunc  {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if validateToken(ctx, token) {
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}


