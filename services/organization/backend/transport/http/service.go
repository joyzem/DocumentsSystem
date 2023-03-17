package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/organization/backend/transport"
	"github.com/joyzem/documents/services/organization/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)
	options = append(options, errorEncoder)

	router.Methods("POST").Path("/organizations").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateOrganization,
			decodeCreateOrganizationRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/organizations").Handler(
		kithttp.NewServer(
			svcEndpoints.GetOrganizations,
			decodeGetOrganizationsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/organizations").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateOrganization,
			decodeUpdateOrganizationRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/organizations").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteOrganization,
			decodeDeleteOrganizationRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/organizations/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.OrganizationById,
			decodeOrganizationByIdRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateOrganizationRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetOrganizationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return dto.GetOrganizationsRequest{}, nil
}

func decodeUpdateOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateOrganizationRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteOrganizationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteOrganizationRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeOrganizationByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		base.LogError(err)
	}
	return dto.OrganizationByIdRequest{Id: id}, err
}
