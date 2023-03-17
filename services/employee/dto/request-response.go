package dto

import "github.com/joyzem/documents/services/employee/domain"

type CreateEmployeeRequest struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	MiddleName          string `json:"middle_name"`
	Post                string `json:"post"`
	PassportSeries      string `json:"passport_series"`
	PassportNumber      string `json:"passport_number"`
	PassportIssuedBy    string `json:"passport_issued_by"`
	PassportDateOfIssue string `json:"passport_date_of_issue"`
}

type CreateEmployeeResponse struct {
	Employee *domain.Employee `json:"employee,omitempty"`
	Err      string           `json:"error,omitempty"`
}

type GetEmployeesRequest struct {
}

type GetEmployeesResponse struct {
	Employees []domain.Employee `json:"employees,omitempty"`
	Err       string            `json:"error,omitempty"`
}

type EmployeeByIdRequest struct {
	Id int `json:"id"`
}

type EmployeeByIdResponse struct {
	Employee *domain.Employee `json:"employee,omitempty"`
	Err      string           `json:"error,omitempty"`
}

type DeleteEmployeeRequest struct {
	Id int `json:"id"`
}

type DeleteEmployeeResponse struct {
	Err string `json:"error,omitempty"`
}

type UpdateEmployeeRequest struct {
	Employee domain.Employee `json:"employee"`
}

type UpdateEmployeeResponse struct {
	Employee *domain.Employee `json:"employee,omitempty"`
	Err      string           `json:"error,omitempty"`
}
