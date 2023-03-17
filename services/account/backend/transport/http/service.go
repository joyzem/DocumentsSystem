package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/account/backend/transport"
	"github.com/joyzem/documents/services/account/dto"
	"github.com/joyzem/documents/services/base"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {

	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)

	options = append(options, errorEncoder)

	router.Methods("POST").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateAccount,
			decodeCreateAccountRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.GetAccounts,
			decodeGetAccountsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateAccount,
			decodeUpdateAccountRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteAccount,
			decodeDeleteAccountRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/accounts/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.AccountById,
			decodeAccountByIdRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateAccountRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetAccountsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return dto.GetAccountsRequest{}, nil
}

func decodeDeleteAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteAccountRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeUpdateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateAccountRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeAccountByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	req := dto.AccountByIdRequest{Id: id}
	return req, nil
}
