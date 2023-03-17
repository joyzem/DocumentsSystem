package service

import "github.com/joyzem/documents/services/employee/domain"

type EmployeeService interface {
	CreateEmployee(
		firstName string,
		lastName string,
		middleName string,
		post string,
		passportSeries string,
		passportNumber string,
		passportIssuedBy string,
		passportDateOfIssue string,
	) (*domain.Employee, error)
	GetEmployees() ([]domain.Employee, error)
	EmployeeById(id int) (*domain.Employee, error)
	DeleteEmployee(id int) error
	UpdateEmployee(employee domain.Employee) (*domain.Employee, error)
}
