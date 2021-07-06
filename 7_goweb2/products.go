package main

type Products struct {
	Products []Product
}
type Product struct {
	Id              int    `json:"id" form:"id"`
	Nombre          string `json:"nombre" form:"nombre" binding:"required"`
	Color           string `json:"color" form:"color" binding:"required"`
	Precio          int    `json:"precio" form:"precio" binding:"required"`
	Stock           int    `json:"stock" form:"stock" binding:"required"`
	Codigo          string `json:"codigo" form:"codigo" binding:"required"`
	Publicado       bool   `json:"publicado" form:"publicado"`
	FechaDecreacion string `json:"fechaCreacion" form:"fechaCreacion"`
}
