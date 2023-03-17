package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/employee/backend/transport"
	"github.com/joyzem/documents/services/employee/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)
	options = append(options, errorEncoder)

	router.Methods("POST").Path("/employees").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateEmployee,
			decodeCreateEmployeeRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/employees").Handler(
		kithttp.NewServer(
			svcEndpoints.GetEmployees,
			decodeGetEmployeesRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/employees").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateEmployee,
			decodeUpdateEmployeeRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/employees").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteEmployee,
			decodeDeleteEmployeeRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/employees/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.EmployeeById,
			decodeEmployeeByIdRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateEmployeeRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetEmployeesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return dto.GetEmployeesRequest{}, nil
}

func decodeUpdateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdateEmployeeRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeleteEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeleteEmployeeRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeEmployeeByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		base.LogError(err)
	}
	return dto.EmployeeByIdRequest{Id: id}, err
}
