package utils

import (
	"fmt"

	customerDto "github.com/joyzem/documents/services/customer/dto"
	"github.com/levigross/grequests"
)

func GetCustomers() (customerDto.GetCustomersResponse, error) {
	customersUrl := fmt.Sprintf("%s/customers", GetCustomersAddress())
	customerResp, err := grequests.Get(customersUrl, nil)
	var customers customerDto.GetCustomersResponse
	customerResp.JSON(&customers)
	return customers, err
}

func GetCustomerById(id int) (customerDto.CustomerByIdResponse, error) {
	customerUrl := fmt.Sprintf("%s/customers/%d", GetCustomersAddress(), id)
	customerResp, err := grequests.Get(customerUrl, &grequests.RequestOptions{
		JSON: customerDto.CustomerByIdRequest{
			Id: id,
		}})
	var customer customerDto.CustomerByIdResponse
	customerResp.JSON(&customer)
	return customer, err
}
