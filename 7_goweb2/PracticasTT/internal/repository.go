package internal

import (
    "fmt"
)

type Product struct {
    Id            int     `form:"id" json:"id"`
    Nombre        string  `form:"nombre" json:"nombre"`
    Color         string  `form:"color" json:"color"`
    Precio        float64 `form:"precio" json:"precio"`
    Stock         uint    `form:"stock" json:"stock"`
    Codigo        string  `form:"codigo" json:"codigo"`
    Publicado     bool    `form:"publicado" json:"publicado"`
    FechaCreacion string  `form:"fechacreacion" json:"fechacreacion"`
}

//Variable Global
var Products []Product

type Repository interface {
    FilteredProducts(nombre string) []Product
    GetAll() []Product
    GetById(id int) (Product, error)
    CreateProduct(nombre string, color string, precio float64, stock uint, codigo string, publicado bool, fechaCreacion string) Product
}

type repository struct{}

func NewRepository() Repository {
    return &repository{}
}

func (r *repository) FilteredProducts(nombre string) []Product {
    var filtrados []Product
    cantidadFiltrados := 0
    for _, p := range Products {
        if nombre == p.Nombre {
            filtrados = append(filtrados, p)
            cantidadFiltrados++
        }
    }
    return filtrados
}

func (r *repository) GetAll() []Product {
    return Products
}

func (r *repository) GetById(id int) (Product, error) {
    var producto Product

    if id != 0 {
        for _, p := range Products {
            if id == p.Id {
                producto = p
            }
        }
        return producto, nil
    } else {
        return producto, fmt.Errorf("el ID %d no es valido", id)
    }
}

func (r *repository) CreateProduct(nombre string, color string, precio float64, stock uint, codigo string, publicado bool, fechaCreacion string) Product {
    product := Product{
        Nombre:        nombre,
        Color:         color,
        Precio:        precio,
        Publicado:     publicado,
        FechaCreacion: fechaCreacion,
    }

    var id int
    for _, i := range Products {
        id = i.Id
        id++
    }
    if id < 1 {
        id = 1
    }
    product.Id = id

    Products = append(Products, product)

    return product
}