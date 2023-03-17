package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/backend/transport"
	"github.com/joyzem/documents/services/proxy/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)
	options = append(options, errorEncoder)
	
	router.Methods("POST").Path("/proxy").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProxyHeader,
			decodeCreateProxyHeaderRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/proxy").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProxies,
			decodeGetProxiesRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/proxy/{id:[0-9]+}").Handler(
		kithttp.NewServer(
			svcEndpoints.ProxyById,
			decodeProxyByIdRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/proxy").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateProxyHeader,
			decodeUpdateProxyHeaderRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/proxy").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProxy,
			decodeDeleteProxyRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("POST").Path("/proxy/body").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProxyBodyItem,
			decodeCreateProxyBodyItemRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/proxy/body").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProxyBodyItem,
			decodeDeleteProxyBodyItemRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateProxyHeaderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateProxyHeaderRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetProxiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetProxiesRequest
	return req, nil
}

func decodeProxyByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	return dto.ProxyByIdRequest{Id: id}, err
}

func decodeUpdateProxyHeaderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateProxyHeaderRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteProxyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteProxyRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeCreateProxyBodyItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateProxyBodyItemRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteProxyBodyItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteProxyBodyItemRequest
	err := base.DecodeBody(r, &req)
	return req, err
}
