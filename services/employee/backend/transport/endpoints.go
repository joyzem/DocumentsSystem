package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/employee/backend/service"
	"github.com/joyzem/documents/services/employee/dto"
)

type Endpoints struct {
	CreateEmployee endpoint.Endpoint
	GetEmployees   endpoint.Endpoint
	EmployeeById   endpoint.Endpoint
	DeleteEmployee endpoint.Endpoint
	UpdateEmployee endpoint.Endpoint
}

func MakeEndpoints(s service.EmployeeService) Endpoints {
	return Endpoints{
		CreateEmployee: makeCreateEmployeeEndpoint(s),
		GetEmployees:   makeGetEmployeesEndpoint(s),
		EmployeeById:   makeEmployeeByIdEndpoint(s),
		DeleteEmployee: makeDeleteEmployeeEndpoint(s),
		UpdateEmployee: makeUpdateEmployeeEndpoint(s),
	}
}

func makeCreateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateEmployeeRequest)
		employee, err := s.CreateEmployee(
			req.FirstName,
			req.LastName,
			req.MiddleName,
			req.Post,
			req.PassportSeries,
			req.PassportNumber,
			req.PassportIssuedBy,
			req.PassportDateOfIssue,
		)
		return dto.CreateEmployeeResponse{Employee: employee}, err
	}
}

func makeGetEmployeesEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		employees, err := s.GetEmployees()
		return dto.GetEmployeesResponse{Employees: employees}, err
	}
}

func makeEmployeeByIdEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.EmployeeByIdRequest)
		employee, err := s.EmployeeById(req.Id)
		return dto.EmployeeByIdResponse{Employee: employee}, err
	}
}

func makeDeleteEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, e error) {
		req := request.(dto.DeleteEmployeeRequest)
		err := s.DeleteEmployee(req.Id)
		return dto.DeleteEmployeeResponse{}, err
	}
}

func makeUpdateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.UpdateEmployeeRequest)
		employee, err := s.UpdateEmployee(req.Employee)
		return dto.UpdateEmployeeResponse{Employee: employee}, err
	}
}
