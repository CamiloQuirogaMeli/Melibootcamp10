package internal
type ServiceInterface interface {
    GetAll()
    GetByFilter()
    GetById()
    CreateProduct(nombre string, color string, precio float64, stock uint, codigo string, publicado bool, fechaCreacion string) Product
    NewService() *Service
}

type Service struct {
    repository Repository
}

func NewService(r Repository) *Service {
    return &Service{
        repository: r,
    }
}

func (s *Service) GetAll() []Product {
    return s.repository.GetAll()
}

func (s *Service) GetByFilter(nombre string) []Product {
    return s.repository.FilteredProducts(nombre)

}

func (s *Service) GetById(id int) (Product, error) {
    return s.repository.GetById(id)
}

func (s *Service) CreateProduct(nombre string, color string, precio float64, stock uint, codigo string, publicado bool, fechaCreacion string) Product {
    return s.repository.CreateProduct(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}