package products

type Repository interface {
	getAll() ([]Producto, error)
	store(nombre string, cantidad int, precios float64) (Producto, error)
}

type Producto struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}
