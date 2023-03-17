package repo

import "github.com/joyzem/documents/services/employee/domain"

type EmployeeRepo interface {
	CreateEmployee(domain.Employee) (*domain.Employee, error)
	GetEmployees() ([]domain.Employee, error)
	EmployeeById(id int) (*domain.Employee, error)
	DeleteEmployee(id int) error
	UpdateEmployee(domain.Employee) (*domain.Employee, error)
}
