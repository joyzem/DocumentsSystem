package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/customer/domain"
	"github.com/joyzem/documents/services/customer/dto"
	"github.com/joyzem/documents/services/customer/frontend/utils"
	"github.com/levigross/grequests"
)

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/customers", utils.GetBackendAddress())
	resp, _ := grequests.Get(url, nil)
	var data dto.GetCustomersResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
		return
	}
	customersPage, _ := template.ParseFiles("../static/html/customers.html")
	customersPage.Execute(w, data.Customers)
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	body := dto.DeleteCustomerRequest{Id: id}
	customersUrl := fmt.Sprintf("%s/customers", utils.GetBackendAddress())
	resp, _ := grequests.Delete(customersUrl, &grequests.RequestOptions{
		JSON: body,
	})
	var deleteResponse dto.DeleteCustomerResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/customer/customers", http.StatusSeeOther)
}

func CreateCustomerGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/create-customer.html")
}

func CreateCustomerPostHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	request := dto.CreateCustomerRequest{
		Name: name,
	}
	customersUrl := fmt.Sprintf("%s/customers", utils.GetBackendAddress())
	resp, _ := grequests.Post(customersUrl, &grequests.RequestOptions{
		JSON: request,
	})
	var data dto.CreateCustomerResponse
	resp.JSON(&data)
	if data.Err == "" {
		http.Redirect(w, r, "/customer/customers", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
}

func UpdateCustomerGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	url := fmt.Sprintf("%s/customers/%d", utils.GetBackendAddress(), id)
	resp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.CustomerByIdRequest{Id: id},
	})
	var data dto.CustomerByIdResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusBadRequest)
		return
	}
	updateCustomerPage, _ := template.ParseFiles("../static/html/update-customer.html")
	updateCustomerPage.Execute(w, data.Customer)
}

func UpdateCustomerPostHandler(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	customer := domain.Customer{
		Id:   customerId,
		Name: name,
	}
	request := dto.UpdateCustomerRequest{
		Customer: customer,
	}
	customersUrl := fmt.Sprintf("%s/customers", utils.GetBackendAddress())
	resp, _ := grequests.Put(customersUrl, &grequests.RequestOptions{
		JSON: request,
	})
	var data dto.UpdateCustomerResponse
	resp.JSON(&data)
	if data.Err == "" {
		http.Redirect(w, r, "/customer/customers", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
}
