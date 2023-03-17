package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/customer/backend/service"
	"github.com/joyzem/documents/services/customer/dto"
)

type Endpoints struct {
	CreateCustomer endpoint.Endpoint
	GetCustomers   endpoint.Endpoint
	UpdateCustomer endpoint.Endpoint
	DeleteCustomer endpoint.Endpoint
	CustomerById   endpoint.Endpoint
}

func MakeEndpoints(s service.CustomerService) Endpoints {
	return Endpoints{
		CreateCustomer: makeCreateCustomerEndpoint(s),
		GetCustomers:   makeGetCustomersEndpoint(s),
		UpdateCustomer: makeUpdateCustomerEndpoint(s),
		DeleteCustomer: makeDeleteCustomerEndpoint(s),
		CustomerById:   makeCustomerByIdEndpoint(s),
	}
}

func makeCreateCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateCustomerRequest)
		customer, err := s.CreateCustomer(req.Name)
		return dto.CreateCustomerResponse{Customer: customer}, err
	}
}

func makeGetCustomersEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		customers, err := s.GetCustomers()
		return dto.GetCustomersResponse{Customers: customers}, err
	}
}

func makeUpdateCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateCustomerRequest)
		customer, err := s.UpdateCustomer(req.Customer)
		return dto.UpdateCustomerResponse{Customer: customer}, err
	}
}

func makeDeleteCustomerEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteCustomerRequest)
		return dto.DeleteCustomerResponse{}, s.DeleteCustomer(req.Id)
	}
}

func makeCustomerByIdEndpoint(s service.CustomerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CustomerByIdRequest)
		customer, err := s.CustomerById(req.Id)
		return dto.CustomerByIdResponse{Customer: customer}, err
	}
}
