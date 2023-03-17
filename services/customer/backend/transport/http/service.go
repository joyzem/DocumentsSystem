package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/customer/backend/transport"
	"github.com/joyzem/documents/services/customer/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)
	options = append(options, errorEncoder)

	router.Methods("POST").Path("/customers").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateCustomer,
			decodeCreateCustomerRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/customers").Handler(
		kithttp.NewServer(
			svcEndpoints.GetCustomers,
			decodeGetCustomersRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/customers").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateCustomer,
			decodeUpdateCustomerRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/customers").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteCustomer,
			decodeDeleteCustomerRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/customers/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.CustomerById,
			decodeCustomerByIdRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateCustomerRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetCustomersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// У этого запроса нет тела, и декодировать нечего
	return dto.GetCustomersRequest{}, nil
}

func decodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateCustomerRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteCustomerRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	req := dto.CustomerByIdRequest{Id: id}
	return req, nil
}
