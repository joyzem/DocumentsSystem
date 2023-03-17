package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/proxy/backend/service"
	"github.com/joyzem/documents/services/proxy/dto"
)

type Endpoints struct {
	CreateProxyHeader   endpoint.Endpoint
	CreateProxyBodyItem endpoint.Endpoint
	GetProxies          endpoint.Endpoint
	ProxyById           endpoint.Endpoint
	UpdateProxyHeader   endpoint.Endpoint
	DeleteProxy         endpoint.Endpoint
	DeleteProxyBodyItem endpoint.Endpoint
}

func MakeEndpoints(s service.ProxyService) Endpoints {
	return Endpoints{
		CreateProxyHeader:   makeCreateProxyHeaderEndpoint(s),
		CreateProxyBodyItem: makeCreateProxyBodyItemEndpoint(s),
		GetProxies:          makeGetProxiesEndpoint(s),
		ProxyById:           makeProxyByIdEndpoint(s),
		UpdateProxyHeader:   makeUpdateProxyHeaderEndpoint(s),
		DeleteProxy:         makeDeleteProxyEndpoint(s),
		DeleteProxyBodyItem: makeDeleteProxyBodyItemEndpoint(s),
	}
}

func makeCreateProxyHeaderEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateProxyHeaderRequest)
		header, err := s.CreateProxyHeader(
			req.OrganizationId,
			req.CustomerId,
			req.EmployeeId,
			req.DateOfIssue,
			req.IsValidUntil,
		)
		return dto.CreateProxyHeaderResponse{ProxyHeader: header}, err
	}
}

func makeCreateProxyBodyItemEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreateProxyBodyItemRequest)
		proxy, err := s.CreateProxyBodyItem(req.ProductId, req.ProxyId, req.ProductAmount)
		return dto.CreateProxyBodyItemResponse{Proxy: proxy}, err
	}
}

func makeGetProxiesEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		proxies, err := s.GetProxies()
		return dto.GetProxiesResponse{Proxies: proxies}, err
	}
}

func makeProxyByIdEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.ProxyByIdRequest)
		proxy, err := s.ProxyById(req.Id)
		return dto.ProxyByIdResponse{Proxy: proxy}, err
	}
}

func makeUpdateProxyHeaderEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.UpdateProxyHeaderRequest)
		proxy, err := s.UpdateProxyHeader(req.Header)
		return dto.UpdateProxyResponse{Proxy: proxy}, err
	}
}

func makeDeleteProxyBodyItemEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DeleteProxyBodyItemRequest)
		err = s.DeleteProxyBodyItem(req.Id)
		return dto.DeleteProxyResponse{}, err
	}
}

func makeDeleteProxyEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DeleteProxyRequest)
		err = s.DeleteProxy(req.Id)
		return dto.DeleteProxyResponse{}, err
	}
}
