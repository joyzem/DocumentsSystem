package service

import "github.com/joyzem/documents/services/proxy/domain"

type ProxyService interface {
	CreateProxyHeader(organizationId int, customerId int, employeeId int, dateOfIssue string, isValidUntil string) (*domain.ProxyHeader, error)
	CreateProxyBodyItem(productId int, proxyId int, productAmount int) (*domain.Proxy, error)
	GetProxies() ([]domain.ProxyHeader, error)
	ProxyById(id int) (*domain.Proxy, error)
	UpdateProxyHeader(domain.ProxyHeader) (*domain.Proxy, error)
	DeleteProxy(id int) error
	DeleteProxyBodyItem(itemId int) error
}
