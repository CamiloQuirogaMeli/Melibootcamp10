package handler

import (
    "net/http"
    "strconv"

    products "github.com/CamiloQuirogaMeli/7_goweb2/PracticasTT/internal"
    "github.com/gin-gonic/gin"
)

type request struct {
    Nombre        string  `form:"nombre" json:"nombre"`
    Color         string  `form:"color" json:"color"`
    Precio        float64 `form:"precio" json:"precio"`
    Stock         uint    `form:"stock" json:"stock"`
    Codigo        string  `form:"codigo" json:"codigo"`
    Publicado     bool    `form:"publicado" json:"publicado"`
    FechaCreacion string  `form:"fecha_creacion" json:"fecha_creacion"`
}

type Product struct {
    service products.Service
}

func NewProduct(p products.Service) *Product {
    return &Product{p}
}

func (p *Product) Store() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.Request.Header.Get("token")
        if token != "123456" {
            ctx.JSON(401, gin.H{"error": "token inválido"})
            return
        }
        var req request
        if err := ctx.Bind(&req); err != nil {
            ctx.JSON(404, gin.H{
                "error": err.Error(),
            })
            return
        }
        if req.Nombre == "" {
            ctx.JSON(http.StatusOK, gin.H{"error": "Nombre es un campo requerido!"})
            return
        }
        if req.Color == "" {
            ctx.JSON(http.StatusOK, gin.H{"error": "Color es un campo requerido!"})
            return
        }
        if req.Precio == 0 {
            ctx.JSON(http.StatusOK, gin.H{"error": "El precio del producto no puede ser 0!"})
            return
        }
        if req.Stock == 0 {
            ctx.JSON(http.StatusOK, gin.H{"error": "El stock del producto no puede ser 0!"})
            return
        }
        if req.Codigo == "" {
            ctx.JSON(http.StatusOK, gin.H{"error": "Codigo es un campo requerido!"})
            return
        }
        if req.FechaCreacion == "" {
            ctx.JSON(http.StatusOK, gin.H{"error": "La fecha de creacion es un campo requerido!"})
            return
        }

        prod := p.service.CreateProduct(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)

        ctx.JSON(http.StatusCreated, prod)
    }
}

func (p *Product) GetAll() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        if ctx.Query("nombre") != "" {
            prods := p.service.GetByFilter(ctx.Query("nombre"))
            ctx.JSON(http.StatusOK, prods)
            return
        }

        productos := p.service.GetAll()

        ctx.JSON(http.StatusOK, productos)
    }
}

func (p *Product) GetById() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        id := ctx.Param("id")
        proId, _ := strconv.Atoi(id)

        producto, err := p.service.GetById(proId)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(http.StatusOK, producto)

    }
}