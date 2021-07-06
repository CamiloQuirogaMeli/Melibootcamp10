package main

type Users struct {
	Users []User
}
type User struct {
	Id              int
	Nombre          string
	Apellido        string
	Email           string
	Edad            int
	Altura          string
	Activo          bool
	FechaDecreacion string
}
