package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/organization/domain"
	"github.com/joyzem/documents/services/organization/dto"
	"github.com/joyzem/documents/services/organization/frontend/utils"

	accountsDomain "github.com/joyzem/documents/services/account/domain"
	accountsDto "github.com/joyzem/documents/services/account/dto"
	"github.com/levigross/grequests"
)

func OrganizationsHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/organizations", utils.GetOrganizationsAddress())
	resp, _ := grequests.Get(url, nil)
	var data dto.GetOrganizationsResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
		return
	}
	tmpl, _ := template.ParseFiles("../static/html/organizations.html")
	tmpl.Execute(w, data.Organizations)
}

func DeleteOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	body := dto.DeleteOrganizationRequest{Id: id}
	url := fmt.Sprintf("%s/organizations", utils.GetOrganizationsAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: body,
	})
	var deleteResponse dto.DeleteOrganizationResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/organization/organizations", http.StatusSeeOther)
}

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetAccountsAddress())
	resp, _ := grequests.Get(accountsUrl, nil)
	var data accountsDto.GetAccountsResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
	tmpl, _ := template.ParseFiles("../static/html/create-organization.html")
	tmpl.Execute(w, data.Accounts)
}

func CreateOrganizationPostHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	address := r.FormValue("address")
	accountIdStr := r.FormValue("account_id")
	accountId, accountErr := strconv.Atoi(accountIdStr)
	chief := r.FormValue("chief")
	financialChief := r.FormValue("financial_chief")
	if name == "" ||
		address == "" ||
		accountErr != nil ||
		chief == "" ||
		financialChief == "" {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("%s/organizations", utils.GetOrganizationsAddress())
	request := dto.CreateOrganizationRequest{
		Name:           name,
		Address:        address,
		AccountId:      accountId,
		Chief:          chief,
		FinancialChief: financialChief,
	}
	resp, _ := grequests.Post(url, &grequests.RequestOptions{
		JSON: request,
	})
	var data dto.CreateOrganizationResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/organization/organizations", http.StatusSeeOther)
	}
}

func OrganizationDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	organizationUrl := fmt.Sprintf("%s/organizations/%d", utils.GetOrganizationsAddress(), id)
	resp, _ := grequests.Get(organizationUrl, &grequests.RequestOptions{
		JSON: dto.OrganizationByIdRequest{Id: id},
	})
	var organizationResp dto.OrganizationByIdResponse
	resp.JSON(&organizationResp)
	if organizationResp.Err != "" {
		http.Error(w, organizationResp.Err, http.StatusInternalServerError)
		return
	}

	accountsUrl := fmt.Sprintf("%s/accounts/%d", utils.GetAccountsAddress(), organizationResp.Organization.AccountId)
	resp, _ = grequests.Get(accountsUrl, &grequests.RequestOptions{
		JSON: accountsDto.AccountByIdRequest{Id: organizationResp.Organization.AccountId},
	})
	var accountResp accountsDto.AccountByIdResponse
	resp.JSON(&accountResp)
	if accountResp.Err != "" {
		http.Error(w, accountResp.Err, http.StatusInternalServerError)
		return
	}

	type organizationDetail struct {
		Name               string
		Address            string
		BankName           string
		BankIdentityNumber string
		Chief              string
		FinancialChief     string
	}
	tmpl, _ := template.ParseFiles("../static/html/organization-details.html")
	data := organizationDetail{
		Name:               organizationResp.Organization.Name,
		Address:            organizationResp.Organization.Address,
		BankName:           accountResp.Account.BankName,
		BankIdentityNumber: accountResp.Account.BankIdentityNumber,
		Chief:              organizationResp.Organization.Chief,
		FinancialChief:     organizationResp.Organization.FinancialChief,
	}
	tmpl.Execute(w, data)
}

func UpdateOrganizationGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	url := fmt.Sprintf("%s/organizations/%d", utils.GetOrganizationsAddress(), id)
	orgResp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.OrganizationByIdRequest{Id: id},
	})
	var organization dto.OrganizationByIdResponse
	orgResp.JSON(&organization)

	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetAccountsAddress())
	accResp, _ := grequests.Get(accountsUrl, nil)
	var accounts accountsDto.GetAccountsResponse
	accResp.JSON(&accounts)
	if accounts.Err != "" {
		http.Error(w, accounts.Err, http.StatusInternalServerError)
	}

	type templateData struct {
		Organization domain.Organization
		Accounts     []accountsDomain.Account
	}
	data := templateData{
		Organization: *organization.Organization,
		Accounts:     accounts.Accounts,
	}

	tmpl, _ := template.ParseFiles("../static/html/update-organization.html")
	tmpl.Execute(w, data)
}

func UpdateOrganizationPostHandler(w http.ResponseWriter, r *http.Request) {
	id, idErr := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	address := r.FormValue("address")
	accountIdStr := r.FormValue("account_id")
	accountId, accountErr := strconv.Atoi(accountIdStr)
	chief := r.FormValue("chief")
	financialChief := r.FormValue("financial_chief")
	if name == "" ||
		address == "" ||
		accountErr != nil ||
		chief == "" ||
		financialChief == "" ||
		idErr != nil {

		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("%s/organizations", utils.GetOrganizationsAddress())
	organization := domain.Organization{
		Id:             id,
		Name:           name,
		Address:        address,
		AccountId:      accountId,
		Chief:          chief,
		FinancialChief: financialChief,
	}
	request := dto.UpdateOrganizationRequest{
		Organization: organization,
	}
	resp, _ := grequests.Put(url, &grequests.RequestOptions{
		JSON: request,
	})
	var data dto.UpdateOrganizationResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/organization/organizations", http.StatusSeeOther)
	}
}
