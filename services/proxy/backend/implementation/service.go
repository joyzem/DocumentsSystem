package implementation

import (
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/backend/repo"
	"github.com/joyzem/documents/services/proxy/backend/service"
	"github.com/joyzem/documents/services/proxy/domain"
)

type proxyService struct {
	proxyRepo repo.ProxyRepo
}

func NewProxyService(proxyRepo repo.ProxyRepo) service.ProxyService {
	return &proxyService{
		proxyRepo: proxyRepo,
	}
}

func (s *proxyService) CreateProxyHeader(organizationId int, customerId int, employeeId int, dateOfIssue string, isValidUntil string) (*domain.ProxyHeader, error) {
	header := domain.ProxyHeader{
		OrganizationId: organizationId,
		CustomerId:     customerId,
		EmployeeId:     employeeId,
		DateOfIssue:    dateOfIssue,
		IsValidUntil:   isValidUntil,
	}
	createdHeader, err := s.proxyRepo.CreateProxyHeader(header)
	base.LogError(err)
	if err != nil {
		return nil, err
	}
	return createdHeader, nil
}

func (s *proxyService) CreateProxyBodyItem(productId int, proxyId int, productAmount int) (*domain.Proxy, error) {
	item := domain.ProxyBodyItem{
		ProductId:     productId,
		ProxyId:       proxyId,
		ProductAmount: productAmount,
	}
	proxy, err := s.proxyRepo.CreateProxyBodyItem(item)
	base.LogError(err)
	return proxy, err
}

func (s *proxyService) GetProxies() ([]domain.ProxyHeader, error) {
	proxies, err := s.proxyRepo.GetProxies()
	base.LogError(err)
	return proxies, err
}

func (s *proxyService) ProxyById(id int) (*domain.Proxy, error) {
	proxy, err := s.proxyRepo.ProxyById(id)
	base.LogError(err)
	return proxy, err
}

func (s *proxyService) UpdateProxyHeader(header domain.ProxyHeader) (*domain.Proxy, error) {
	proxy, err := s.proxyRepo.UpdateProxyHeader(header)
	base.LogError(err)
	return proxy, err
}

func (s *proxyService) DeleteProxy(id int) error {
	err := s.proxyRepo.DeleteProxy(id)
	base.LogError(err)
	return err
}

func (s *proxyService) DeleteProxyBodyItem(itemId int) error {
	err := s.proxyRepo.DeleteProxyBodyItem(itemId)
	base.LogError(err)
	return err
}
