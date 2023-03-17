package repo

import "github.com/joyzem/documents/services/customer/domain"

type CustomerRepo interface {
	GetCustomers() ([]domain.Customer, error)
	CreateCustomer(domain.Customer) (*domain.Customer, error)
	UpdateCustomer(domain.Customer) (*domain.Customer, error)
	DeleteCustomer(int) error
	CustomerById(int) (*domain.Customer, error)
}
