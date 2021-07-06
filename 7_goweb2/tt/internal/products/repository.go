package products

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps = []Product{
	{
		Id:    0,
		Name:  "TV",
		Type:  "Electro",
		Count: 100,
		Price: 2500.0,
	},
	{
		Id:    1,
		Name:  "Colchon",
		Type:  "Hogar",
		Count: 200,
		Price: 1000.0,
	},
}
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre string, tipo string, cantidad int, precio float64) (Product, error)
	LastId() (int, error)
}

type repository struct {}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastId() (int, error) {
	if len(ps) != 0 {
		lastId = ps[len(ps)-1].Id
	}
	return lastId, nil
}

func (r *repository) Store(id int, nombre string, tipo string, cantidad int, precio float64) (Product, error) {
	p:= Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	lastId = p.Id
	return p, nil
}
func NewRepository() Repository {
	return &repository{}
}
