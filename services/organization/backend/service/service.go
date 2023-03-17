package service

import "github.com/joyzem/documents/services/organization/domain"

type OrganizationService interface {
	GetOrganizations() ([]domain.Organization, error)
	CreateOrganization(name string, address string, accountId int, chief string, financialChief string) (*domain.Organization, error)
	UpdateOrganization(domain.Organization) (*domain.Organization, error)
	DeleteOrganization(int) error
	OrganizationById(int) (*domain.Organization, error)
}
