package main

type Products struct {
	Products []Product
}
type Product struct {
	Id              int
	Nombre          string
	Color           string
	Precio          int
	Stock           int
	Codigo          string
	Publicado       bool
	FechaDecreacion string
}
