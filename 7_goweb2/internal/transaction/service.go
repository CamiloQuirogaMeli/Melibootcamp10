package transaction

type Service interface {
	GetAll() ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	Store(transactionCode string, money int, emitter string, receiver string, date string) (Transaction, error)
	LastID() (int, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	return s.repository.GetAll()
}

func (s *service) GetByID(id int) (Transaction, error) {
	return s.repository.GetByID(id)
}

func (s *service) Store(transactionCode string, money int, emitter string, receiver string, date string) (Transaction, error) {
	return s.repository.Store(transactionCode, money, emitter, receiver, date)
}

func (s *service) LastID() (int, error) {
	return s.repository.LastID()
}
