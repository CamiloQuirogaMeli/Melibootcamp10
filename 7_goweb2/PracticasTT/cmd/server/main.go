package main
import (
    "github.com/extviconace/DigitalHouse/cmd/server/handler"
    products "github.com/extviconace/DigitalHouse/internal"

    "github.com/gin-gonic/gin"
)

func main() {
    repo := products.NewRepository()
    service := products.NewService(repo)
    p := handler.NewProduct(*service)

    r := gin.Default()
    pr := r.Group("/productos")
    pr.POST("/", p.Store())
    pr.GET("/:id", p.GetById())
    pr.GET("/", p.GetAll())

    r.Run()
}