package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/organization/backend/service"
	"github.com/joyzem/documents/services/organization/dto"
)

type Endpoints struct {
	CreateOrganization endpoint.Endpoint
	GetOrganizations   endpoint.Endpoint
	UpdateOrganization endpoint.Endpoint
	DeleteOrganization endpoint.Endpoint
	OrganizationById   endpoint.Endpoint
}

func MakeEndpoints(s service.OrganizationService) Endpoints {
	return Endpoints{
		CreateOrganization: makeCreateOrganizationEndpoint(s),
		GetOrganizations:   makeGetOrganizationsEndpoint(s),
		UpdateOrganization: makeUpdateOrganizationsEndpoint(s),
		DeleteOrganization: makeDeleteOrganizationEndpoint(s),
		OrganizationById:   makeOrganizationByIdEndpoint(s),
	}
}

func makeCreateOrganizationEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateOrganizationRequest)
		organization, err := s.CreateOrganization(req.Name, req.Address, req.AccountId, req.Chief, req.FinancialChief)
		return dto.CreateOrganizationResponse{Organization: organization}, err
	}
}

func makeGetOrganizationsEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		organizations, err := s.GetOrganizations()
		return dto.GetOrganizationsResponse{Organizations: organizations}, err
	}
}

func makeUpdateOrganizationsEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateOrganizationRequest)
		organization, err := s.UpdateOrganization(req.Organization)
		return dto.UpdateOrganizationResponse{Organization: organization}, err
	}
}

func makeDeleteOrganizationEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteOrganizationRequest)
		err := s.DeleteOrganization(req.Id)
		return dto.DeleteOrganizationResponse{}, err
	}
}

func makeOrganizationByIdEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.OrganizationByIdRequest)
		org, err := s.OrganizationById(req.Id)
		return dto.OrganizationByIdResponse{Organization: org}, err
	}
}
