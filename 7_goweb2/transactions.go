package main

type Transactions struct {
	Transactions []Transaction
}
type Transaction struct {
	Id              int
	Codigo          string
	Moneda          string
	Monto           int
	Emisor          string
	Receptor        string
	FechaDecreacion string
}
