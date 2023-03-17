package dto

import "github.com/joyzem/documents/services/organization/domain"

type GetOrganizationsRequest struct {
}

type GetOrganizationsResponse struct {
	Organizations []domain.Organization `json:"organizations,omitempty"`
	Err           string                `json:"error,omitempty"`
}

type CreateOrganizationRequest struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	AccountId      int    `json:"account_id"`
	Chief          string `json:"chief"`
	FinancialChief string `json:"financial_chief"`
}

type CreateOrganizationResponse struct {
	Organization *domain.Organization `json:"organization,omitempty"`
	Err          string               `json:"error,omitempty"`
}

type UpdateOrganizationRequest struct {
	Organization domain.Organization `json:"organization"`
}

type UpdateOrganizationResponse struct {
	Organization *domain.Organization `json:"organization,omitempty"`
	Err          string               `json:"error,omitempty"`
}

type DeleteOrganizationRequest struct {
	Id int `json:"id"`
}

type DeleteOrganizationResponse struct {
	Err string `json:"error,omitempty"`
}

type OrganizationByIdRequest struct {
	Id int `json:"id"`
}

type OrganizationByIdResponse struct {
	Organization *domain.Organization `json:"organization,omitempty"`
	Err          string               `json:"error,omitempty"`
}
