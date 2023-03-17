package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/domain"
	"github.com/joyzem/documents/services/proxy/dto"
	"github.com/joyzem/documents/services/proxy/frontend/utils"
	"github.com/levigross/grequests"

	accountDto "github.com/joyzem/documents/services/account/dto"

	customerDomain "github.com/joyzem/documents/services/customer/domain"
	employeeDomain "github.com/joyzem/documents/services/employee/domain"
	organizationDomain "github.com/joyzem/documents/services/organization/domain"
)

func ProxiesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение доверенностей. В ответе отсутствуют тела доверенностей,
	// но они и не нужны на этом экране
	proxiesUrl := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Get(proxiesUrl, nil)
	var proxies dto.GetProxiesResponse
	resp.JSON(&proxies)
	if proxies.Err != "" {
		http.Error(w, proxies.Err, http.StatusInternalServerError)
		return
	}

	// Создание структуры для HTML-страницы
	type proxyTemplate struct {
		Id           int
		Organization string
		DateOfIssue  string
		IsValidUntil string
		Employee     string
	}

	templateData := []proxyTemplate{}

	// Цикл по всем доверенностям, чтобы заполнить
	// данные по организации и сотруднику
	for _, proxy := range proxies.Proxies {

		organization, _ := utils.GetOrganizationById(proxy.OrganizationId)
		var organizationName string
		if organization.Err != "" {
			organizationName = fmt.Sprintf("Ошибка сервера: %s", organization.Err)
		} else {
			organizationName = organization.Organization.Name
		}

		employee, _ := utils.GetEmployeeById(proxy.EmployeeId)
		var employeeFullname string
		if employee.Err != "" {
			employeeFullname = fmt.Sprintf("Ошибка сервера: %s", employee.Err)
		} else {
			// Функция форматирования имени. Выдает строку с полным ФИО
			employeeFullname = utils.Fullname(
				employee.Employee.LastName,
				employee.Employee.FirstName,
				employee.Employee.MiddleName,
			)
		}

		// Добавление в массив новой доверенности
		proxyTmpl := proxyTemplate{
			Id:           proxy.Id,
			Organization: organizationName,
			DateOfIssue:  proxy.DateOfIssue,
			IsValidUntil: proxy.IsValidUntil,
			Employee:     employeeFullname,
		}
		templateData = append(templateData, proxyTmpl)
	}
	// Выполнение шаблона
	tmpl, _ := template.ParseFiles("../static/html/proxies.html")
	tmpl.Execute(w, templateData)
}

func CreateProxyGetHandler(w http.ResponseWriter, r *http.Request) {

	type proxyTemplate struct {
		Organizations []organizationDomain.Organization
		Employees     []employeeDomain.Employee
		Customers     []customerDomain.Customer
	}

	templateData := proxyTemplate{}

	organizations, _ := utils.GetOrganizations()
	if organizations.Err != "" {
		base.LogError(errors.New(organizations.Err))
	} else {
		templateData.Organizations = organizations.Organizations
	}

	employees, _ := utils.GetEmployees()
	if employees.Err != "" {
		base.LogError(errors.New(employees.Err))
	} else {
		templateData.Employees = employees.Employees
	}

	customers, _ := utils.GetCustomers()
	if customers.Err != "" {
		base.LogError(errors.New(customers.Err))
	} else {
		templateData.Customers = customers.Customers
	}

	tmpl, _ := template.ParseFiles("../static/html/create-proxy.html")
	tmpl.Execute(w, templateData)
}

func CreateProxyPosthandler(w http.ResponseWriter, r *http.Request) {
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	customerId, _ := strconv.Atoi(r.FormValue("customer_id"))
	employeeId, _ := strconv.Atoi(r.FormValue("employee_id"))
	dateOfIssue := r.FormValue("date_of_issue")
	isValidUntil := r.FormValue("is_valid_until")

	proxy := dto.CreateProxyHeaderRequest{
		OrganizationId: organizationId,
		CustomerId:     customerId,
		EmployeeId:     employeeId,
		DateOfIssue:    dateOfIssue,
		IsValidUntil:   isValidUntil,
	}

	proxyUrl := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Post(proxyUrl, &grequests.RequestOptions{
		JSON: proxy,
	})

	var proxyResp dto.CreateProxyHeaderResponse
	resp.JSON(&proxyResp)

	if proxyResp.Err != "" {
		http.Error(w, proxyResp.Err, http.StatusInternalServerError)
		return
	}

	redirectAddress := fmt.Sprintf("/documents/proxies/update/%d", proxyResp.ProxyHeader.Id)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}

func UpdateProxyGetHandler(w http.ResponseWriter, r *http.Request) {
	// Получение доверенности
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	// Создание шаблонов для HTML-страницы
	type proxyBodyTemplate struct {
		Id            int
		Name          string
		Price         int
		ProductAmount int
	}

	type updateProxyTemplate struct {
		Proxy          domain.Proxy
		Organizations  []organizationDomain.Organization
		Employees      []employeeDomain.Employee
		Customers      []customerDomain.Customer
		ProxyBodyItems []proxyBodyTemplate
	}

	templateData := updateProxyTemplate{
		Proxy: *proxy.Proxy,
	}

	// Цикл по товарам из доверенности, чтобы
	// заполнить тело доверенности
	proxyBodyItems := []proxyBodyTemplate{}
	for _, bodyItem := range proxy.Proxy.ProxyBodyItems {
		// Получение товара
		product, _ := utils.GetProductById(bodyItem.ProductId)
		if product.Err != "" {
			base.LogError(errors.New(product.Err))
		} else {
			// Добавление в список товаров
			proxyBodyItems = append(proxyBodyItems, proxyBodyTemplate{
				Id:            product.Product.Id,
				Name:          product.Product.Name,
				Price:         product.Product.Price,
				ProductAmount: bodyItem.ProductAmount,
			})
		}
	}
	// Инициализация тела доверенности
	templateData.ProxyBodyItems = proxyBodyItems

	// Получение организаций и инициализация в шаблоне
	organizations, _ := utils.GetOrganizations()
	if organizations.Err != "" {
		base.LogError(errors.New(organizations.Err))
	} else {
		templateData.Organizations = organizations.Organizations
	}

	// Инициализация аналогично организациям
	employees, _ := utils.GetEmployees()
	if employees.Err != "" {
		base.LogError(errors.New(employees.Err))
	} else {
		templateData.Employees = employees.Employees
	}

	// Инициализация контрагентов
	customers, _ := utils.GetCustomers()
	if customers.Err != "" {
		base.LogError(errors.New(customers.Err))
	} else {
		templateData.Customers = customers.Customers
	}

	// Выполнение шаблона
	tmpl, _ := template.ParseFiles("../static/html/update-proxy.html")
	tmpl.Execute(w, templateData)
}

func UpdateProxyPostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	customerId, _ := strconv.Atoi(r.FormValue("customer_id"))
	employeeId, _ := strconv.Atoi(r.FormValue("employee_id"))
	dateOfIssue := r.FormValue("date_of_issue")
	isValidUntil := r.FormValue("is_valid_until")

	proxy := dto.UpdateProxyHeaderRequest{
		Header: domain.ProxyHeader{
			Id:             id,
			OrganizationId: organizationId,
			CustomerId:     customerId,
			EmployeeId:     employeeId,
			DateOfIssue:    dateOfIssue,
			IsValidUntil:   isValidUntil,
		}}

	proxyUrl := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Put(proxyUrl, &grequests.RequestOptions{
		JSON: proxy,
	})

	var proxyResp dto.UpdateProxyResponse
	resp.JSON(&proxyResp)

	if proxyResp.Err != "" {
		http.Error(w, proxyResp.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/documents/proxies", http.StatusSeeOther)
}

func DeleteProxyHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	body := dto.DeleteProxyRequest{Id: id}
	url := fmt.Sprintf("%s/proxy", utils.GetProxiesAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: body,
	})
	var deleteResponse dto.DeleteProxyResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/documents/proxies", http.StatusSeeOther)
}

func ProxyDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение доверенности
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	proxy, _ := utils.GetProxyById(id)
	if proxy.Err != "" {
		http.Error(w, proxy.Err, http.StatusInternalServerError)
		return
	}

	// Создание структур для шаблона
	type proxyBodyTemplate struct {
		ProductId     int
		Name          string
		Price         int
		ProductAmount int
	}

	type proxyDetailTemplate struct {
		ProxyId             int
		Organization        string
		Address             string
		Account             string
		BankName            string
		BankIdentityNumber  string
		Employee            string
		PassportSeries      string
		PassportNumber      string
		PassportIssuedBy    string
		PassportDateOfIssue string
		Customer            string
		Chief               string
		FinancialChief      string
		DateOfIssue         string
		IsValidUntil        string
		ProxyBodyItems      []proxyBodyTemplate
	}

	// Инициализация данных по доверенности
	templateData := proxyDetailTemplate{
		ProxyId:      proxy.Proxy.ProxyHeader.Id,
		DateOfIssue:  proxy.Proxy.ProxyHeader.DateOfIssue,
		IsValidUntil: proxy.Proxy.ProxyHeader.IsValidUntil,
	}

	// Заполнение списка товаров
	proxyBodyItems := []proxyBodyTemplate{}
	for _, bodyItem := range proxy.Proxy.ProxyBodyItems {
		product, _ := utils.GetProductById(bodyItem.ProductId)
		if product.Err != "" {
			base.LogError(errors.New(product.Err))
		} else {
			proxyBodyItems = append(proxyBodyItems, proxyBodyTemplate{
				ProductId:     product.Product.Id,
				Name:          product.Product.Name,
				Price:         product.Product.Price,
				ProductAmount: bodyItem.ProductAmount,
			})
		}
	}
	templateData.ProxyBodyItems = proxyBodyItems

	// Получение данных об организации
	organization, _ := utils.GetOrganizationById(proxy.Proxy.ProxyHeader.OrganizationId)
	if organization.Err != "" {
		base.LogError(errors.New(organization.Err))
		templateData.Organization = fmt.Sprintf("Ошибка: %s", organization.Err)
	} else {
		templateData.Organization = organization.Organization.Name
		templateData.Address = organization.Organization.Address
		// Получение инициалов имени - Фамилия И.О.
		templateData.Chief = utils.NameInitials(organization.Organization.Chief)
		templateData.FinancialChief = utils.NameInitials(organization.Organization.FinancialChief)
	}

	// Получение данных о счёте
	if organization.Organization != nil {
		accountUrl := fmt.Sprintf("%s/accounts/%d", utils.GetAccountsAddress(), organization.Organization.AccountId)
		accountResp, _ := grequests.Get(accountUrl, &grequests.RequestOptions{
			JSON: accountDto.AccountByIdRequest{
				Id: organization.Organization.AccountId,
			}})
		var account accountDto.AccountByIdResponse
		accountResp.JSON(&account)
		if account.Err != "" {
			templateData.Account = fmt.Sprintf("Ошибка: %s", account.Err)
		} else {
			templateData.Account = account.Account.Account
			templateData.BankName = account.Account.BankName
			templateData.BankIdentityNumber = account.Account.BankIdentityNumber
		}
	}

	// Данные о сотруднике, которому выдали доверенность
	employee, _ := utils.GetEmployeeById(proxy.Proxy.ProxyHeader.EmployeeId)
	if employee.Err != "" {
		base.LogError(errors.New(employee.Err))
		templateData.Employee = fmt.Sprintf("Ошибка: %s", employee.Err)
	} else {
		templateData.Employee = utils.Fullname(
			employee.Employee.LastName,
			employee.Employee.FirstName,
			employee.Employee.MiddleName,
		)
		templateData.PassportSeries = employee.Employee.PassportSeries
		templateData.PassportNumber = employee.Employee.PassportNumber
		templateData.PassportDateOfIssue = employee.Employee.PassportDateOfIssue
		templateData.PassportIssuedBy = employee.Employee.PassportIssuedBy
	}

	// Данные о контрагенте
	customer, _ := utils.GetCustomerById(proxy.Proxy.ProxyHeader.CustomerId)
	if customer.Err != "" {
		base.LogError(errors.New(customer.Err))
		templateData.Customer = fmt.Sprintf("Ошибка: %s", customer.Err)
	} else {
		templateData.Customer = customer.Customer.Name
	}

	// Выполнение шаблона
	tmpl, _ := template.ParseFiles("../static/html/proxy-details.html")
	tmpl.Execute(w, templateData)
}
