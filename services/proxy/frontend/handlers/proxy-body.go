package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/dto"
	"github.com/joyzem/documents/services/proxy/frontend/utils"
	"github.com/levigross/grequests"

	productsDomain "github.com/joyzem/documents/services/product/domain"
)

func CreateProxyBodyGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	type proxyDetail struct {
		Id           int
		Organization string
		Employee     string
		Customer     string
		DateOfIssue  string
		IsValidUntil string
	}

	type createProxyBodyTemplate struct {
		Proxy    proxyDetail
		Products []productsDomain.Product
	}

	templateData := createProxyBodyTemplate{
		Proxy: proxyDetail{
			Id:           proxy.Proxy.ProxyHeader.Id,
			DateOfIssue:  proxy.Proxy.ProxyHeader.DateOfIssue,
			IsValidUntil: proxy.Proxy.ProxyHeader.IsValidUntil,
		}}

	products, _ := utils.GetProducts()
	if products.Err != "" {
		base.LogError(errors.New(products.Err))
	} else {
		templateData.Products = products.Products
	}

	organization, _ := utils.GetOrganizationById(proxy.Proxy.ProxyHeader.OrganizationId)
	if organization.Err != "" {
		base.LogError(errors.New(organization.Err))
		templateData.Proxy.Organization = fmt.Sprintf("Ошибка: %s", organization.Err)
	} else {
		templateData.Proxy.Organization = organization.Organization.Name
	}

	employee, _ := utils.GetEmployeeById(proxy.Proxy.ProxyHeader.EmployeeId)
	if employee.Err != "" {
		base.LogError(errors.New(employee.Err))
		templateData.Proxy.Employee = fmt.Sprintf("Ошибка: %s", employee.Err)
	} else {
		templateData.Proxy.Employee = utils.Fullname(
			employee.Employee.LastName,
			employee.Employee.FirstName,
			employee.Employee.MiddleName,
		)
	}

	customer, _ := utils.GetCustomerById(proxy.Proxy.ProxyHeader.CustomerId)
	if customer.Err != "" {
		base.LogError(errors.New(customer.Err))
		templateData.Proxy.Customer = fmt.Sprintf("Ошибка: %s", customer.Err)
	} else {
		templateData.Proxy.Customer = customer.Customer.Name
	}

	tmpl, _ := template.ParseFiles("../static/html/create-proxy-body.html")
	tmpl.Execute(w, templateData)
}

func CreateProxyBodyPostHandler(w http.ResponseWriter, r *http.Request) {
	proxyId, proxyIdErr := strconv.Atoi(r.FormValue("proxy_id"))
	productId, productIdErr := strconv.Atoi(r.FormValue("product_id"))
	productAmount, productAmountErr := strconv.Atoi(r.FormValue("product_amount"))
	if proxyIdErr != nil || productIdErr != nil || productAmountErr != nil {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}

	proxyBodyUrl := fmt.Sprintf("%s/proxy/body", utils.GetProxiesAddress())
	resp, _ := grequests.Post(proxyBodyUrl, &grequests.RequestOptions{
		JSON: dto.CreateProxyBodyItemRequest{
			ProductId:     productId,
			ProxyId:       proxyId,
			ProductAmount: productAmount,
		}})

	var proxy dto.CreateProxyBodyItemResponse
	resp.JSON(&proxy)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	redirectAddress := fmt.Sprintf("/documents/proxies/details/%d", proxyId)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}

func DeleteProxyBodyHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	url := fmt.Sprintf("%s/proxy/body", utils.GetProxiesAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: dto.DeleteProxyBodyItemRequest{
			Id: id,
		}})
	var responseData dto.DeleteProxyResponse
	resp.JSON(&responseData)
	if responseData.Err != "" {
		http.Error(w, responseData.Err, http.StatusInternalServerError)
		return
	}

	proxyId, _ := strconv.Atoi(r.FormValue("proxy_id"))
	redirectAddress := fmt.Sprintf("/documents/proxies/details/%d", proxyId)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}
