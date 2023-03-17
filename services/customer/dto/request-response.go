package dto

import "github.com/joyzem/documents/services/customer/domain"

type CreateCustomerRequest struct {
	Name string `json:"name"`
}

type CreateCustomerResponse struct {
	Customer *domain.Customer `json:"customer,omitempty"`
	Err      string           `json:"error,omitempty"`
}

type GetCustomersRequest struct {
}

type GetCustomersResponse struct {
	Customers []domain.Customer `json:"customers,omitempty"`
	Err       string            `json:"error,omitempty"`
}

type DeleteCustomerRequest struct {
	Id int `json:"id"`
}

type DeleteCustomerResponse struct {
	Err string `json:"error,omitempty"`
}

type UpdateCustomerRequest struct {
	Customer domain.Customer `json:"customer"`
}

type UpdateCustomerResponse struct {
	Customer *domain.Customer `json:"customer,omitempty"`
	Err      string           `json:"error,omitempty"`
}

type CustomerByIdRequest struct {
	Id int `json:"id"`
}

type CustomerByIdResponse struct {
	Customer *domain.Customer `json:"customer,omitempty"`
	Err      string           `json:"error"`
}
