package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/account/domain"
	"github.com/joyzem/documents/services/account/dto"
	"github.com/joyzem/documents/services/account/frontend/utils"
	"github.com/joyzem/documents/services/base"
	"github.com/levigross/grequests"
)

func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	accountsResp, _ := utils.GetAccountsFromBackend()
	if accountsResp.Err != "" {
		http.Error(w, accountsResp.Err, http.StatusInternalServerError)
		return
	}

	accountsPage, _ := template.ParseFiles("../static/html/accounts.html")
	accountsPage.Execute(w, accountsResp.Accounts)
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	body := dto.DeleteAccountRequest{Id: id}
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	resp, _ := grequests.Delete(accountsUrl, &grequests.RequestOptions{
		JSON: body,
	})
	var deleteResponse dto.DeleteAccountResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
}

func CreateAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/create-account.html")
}

func CreateAccountPostHandler(w http.ResponseWriter, r *http.Request) {
	bankName := r.FormValue("bank_name")
	bankIdentityNumber := r.FormValue("bin")
	account := r.FormValue("account")
	if bankName == "" || len(bankIdentityNumber) != 9 {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	request := dto.CreateAccountRequest{
		Account:            account,
		BankName:           bankName,
		BankIdentityNumber: bankIdentityNumber,
	}
	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	resp, _ := grequests.Post(accountsUrl, &grequests.RequestOptions{
		JSON: request,
	})
	var data dto.CreateAccountResponse
	resp.JSON(&data)
	if data.Err == "" {
		http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err, http.StatusInternalServerError)
	}
}

func UpdateAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	url := fmt.Sprintf("%s/accounts/%d", utils.GetBackendAddress(), id)
	resp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.AccountByIdRequest{Id: id},
	})
	var account dto.AccountByIdResponse
	resp.JSON(&account)
	if account.Err != "" {
		http.Error(w, account.Err, http.StatusBadRequest)
		return
	}
	updateAccountPage, _ := template.ParseFiles("../static/html/update-account.html")
	updateAccountPage.Execute(w, account.Account)
}

func UpdateAccountPostHandler(w http.ResponseWriter, r *http.Request) {
	accountId, _ := strconv.Atoi(r.FormValue("id"))
	bankName := r.FormValue("bank_name")
	bankIdentityCode := r.FormValue("bin")
	acc := r.FormValue("account")
	account := domain.Account{
		Id:                 accountId,
		Account:            acc,
		BankName:           bankName,
		BankIdentityNumber: bankIdentityCode,
	}

	accountsUrl := fmt.Sprintf("%s/accounts", utils.GetBackendAddress())
	request := dto.UpdateAccountRequest{
		Account: account,
	}
	response, _ := grequests.Put(accountsUrl, &grequests.RequestOptions{
		JSON: request,
	})

	var updateResponse dto.UpdateAccountResponse
	response.JSON(&request)
	if updateResponse.Err != "" {
		http.Error(w, updateResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/account/accounts", http.StatusSeeOther)
}
