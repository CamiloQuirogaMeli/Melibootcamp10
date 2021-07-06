package products

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil{
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	lastId++
	producto, err := s.repository.Store(lastId, nombre, tipo, cantidad, precio)
	if err != nil{
		return Product{}, err
	}
	return producto, nil
}

func NewService( r Repository) Service {
	return &service{
		repository: r,
	}
}