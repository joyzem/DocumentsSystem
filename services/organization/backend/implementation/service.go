package implementation

import (
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/organization/backend/repo"
	svc "github.com/joyzem/documents/services/organization/backend/service"
	"github.com/joyzem/documents/services/organization/domain"
)

type service struct {
	organizationRepo repo.OrganizationRepo
}

func NewService(organizationRepo repo.OrganizationRepo) svc.OrganizationService {
	return &service{
		organizationRepo: organizationRepo,
	}
}

func (s *service) GetOrganizations() ([]domain.Organization, error) {
	organizations, err := s.organizationRepo.GetOrganizations()
	base.LogError(err)
	return organizations, err
}

func (s *service) CreateOrganization(name string, address string, accountId int, chief string, financialChief string) (*domain.Organization, error) {
	org := domain.Organization{Name: name, Address: address, AccountId: accountId, Chief: chief, FinancialChief: financialChief}
	createdOrganization, err := s.organizationRepo.CreateOrganization(org)
	base.LogError(err)
	return createdOrganization, err
}

func (s *service) UpdateOrganization(newOrganization domain.Organization) (*domain.Organization, error) {
	updatedOrganization, err := s.organizationRepo.UpdateOrganization(newOrganization)
	base.LogError(err)
	return updatedOrganization, err
}

func (s *service) DeleteOrganization(id int) error {
	err := s.organizationRepo.DeleteOrganization(id)
	base.LogError(err)
	return err
}

func (s *service) OrganizationById(id int) (*domain.Organization, error) {
	org, err := s.organizationRepo.OrganizationById(id)
	base.LogError(err)
	return org, err
}
