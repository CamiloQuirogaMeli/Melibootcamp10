package products

type Service interface {
	getAll() ([]Producto, error)
	store(nombre string, cantidad int, precio float64) (Producto, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Producto, error) {
	return s.repository.getAll()
}

func (s *service) store(nombre string, cantidad int, precio float64) (Producto, error) {
	return s.repository.store(nombre, cantidad, precio)
}
