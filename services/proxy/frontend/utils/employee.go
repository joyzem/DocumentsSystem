package utils

import (
	"fmt"

	employeeDto "github.com/joyzem/documents/services/employee/dto"
	"github.com/levigross/grequests"
)

func GetEmployeeById(id int) (*employeeDto.EmployeeByIdResponse, error) {

	employeeUrl := fmt.Sprintf("%s/employees/%d", GetEmployeesAddress(), id)
	employeeResp, err := grequests.Get(employeeUrl, &grequests.RequestOptions{
		JSON: employeeDto.EmployeeByIdRequest{
			Id: id,
		}})
	var employee employeeDto.EmployeeByIdResponse
	employeeResp.JSON(&employee)
	return &employee, err
}

func GetEmployees() (employeeDto.GetEmployeesResponse, error) {
	employeesUrl := fmt.Sprintf("%s/employees", GetEmployeesAddress())
	emplResp, err := grequests.Get(employeesUrl, nil)
	var employees employeeDto.GetEmployeesResponse
	emplResp.JSON(&employees)
	return employees, err
}
