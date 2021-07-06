package main
import (
    "github.com/CamiloQuirogaMeli/7_goweb2/PracticasTT/cmd/server/handler"
    products "github.com/CamiloQuirogaMeli/7_goweb2/PracticasTT/internal"

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