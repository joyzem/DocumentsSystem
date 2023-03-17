package dto

import "github.com/joyzem/documents/services/proxy/domain"

type CreateProxyHeaderRequest struct {
	OrganizationId int    `json:"organization_id"`
	CustomerId     int    `json:"customer_id"`
	EmployeeId     int    `json:"employee_id"`
	DateOfIssue    string `json:"date_of_issue"`
	IsValidUntil   string `json:"is_valid_until"`
}

type CreateProxyHeaderResponse struct {
	ProxyHeader *domain.ProxyHeader `json:"proxy_header,omitempty"`
	Err         string              `json:"error,omitempty"`
}

type CreateProxyBodyItemRequest struct {
	ProductId     int `json:"product_id"`
	ProxyId       int `json:"proxy_id"`
	ProductAmount int `json:"product_amount"`
}

type CreateProxyBodyItemResponse struct {
	Proxy *domain.Proxy `json:"proxy,omitempty"`
	Err   string        `json:"error,omitempty"`
}

type GetProxiesRequest struct{}

type GetProxiesResponse struct {
	Proxies []domain.ProxyHeader `json:"proxies,omitempty"`
	Err     string               `json:"error,omitempty"`
}

type ProxyByIdRequest struct {
	Id int `json:"id"`
}

type ProxyByIdResponse struct {
	Proxy *domain.Proxy `json:"proxy,omitempty"`
	Err   string        `json:"error,omitempty"`
}

type UpdateProxyHeaderRequest struct {
	Header domain.ProxyHeader `json:"proxy_header"`
}

type UpdateProxyResponse struct {
	Proxy *domain.Proxy `json:"proxy,omitempty"`
	Err   string        `json:"error,omitempty"`
}

type DeleteProxyBodyItemRequest struct {
	Id int `json:"id"`
}

type DeleteProxyRequest struct {
	Id int `json:"id"`
}

type DeleteProxyResponse struct {
	Err string `json:"error,omitempty"`
}
