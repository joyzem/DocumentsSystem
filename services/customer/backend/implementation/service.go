package implementation

import (
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/customer/backend/repo"
	svc "github.com/joyzem/documents/services/customer/backend/service"
	"github.com/joyzem/documents/services/customer/domain"
)

type service struct {
	repo repo.CustomerRepo
}

func NewService(repo repo.CustomerRepo) svc.CustomerService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCustomer(name string) (*domain.Customer, error) {
	customer := domain.Customer{
		Name: name,
	}
	result, err := s.repo.CreateCustomer(customer)
	base.LogError(err)
	return result, err
}

func (s *service) GetCustomers() ([]domain.Customer, error) {
	result, err := s.repo.GetCustomers()
	base.LogError(err)
	return result, err
}

func (s *service) UpdateCustomer(customer domain.Customer) (*domain.Customer, error) {
	result, err := s.repo.UpdateCustomer(customer)
	base.LogError(err)
	return result, err
}

func (s *service) DeleteCustomer(id int) error {
	err := s.repo.DeleteCustomer(id)
	base.LogError(err)
	return err
}

func (s *service) CustomerById(id int) (*domain.Customer, error) {
	customer, err := s.repo.CustomerById(id)
	base.LogError(err)
	return customer, err
}
