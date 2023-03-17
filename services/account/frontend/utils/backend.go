package utils

import (
	"encoding/json"
	"fmt"

	"github.com/joyzem/documents/services/account/dto"
	"github.com/joyzem/documents/services/base"
	"github.com/levigross/grequests"
)

const (
	BACKEND_ENV = "ACCOUNT_BACKEND_ADDRESS"
)

func GetBackendAddress() string {
	address := base.GetEnv(BACKEND_ENV, "http://localhost:7073")
	return address
}

func GetAccountsFromBackend() (*dto.GetAccountsResponse, error) {
	url := fmt.Sprintf("%s/accounts", GetBackendAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data dto.GetAccountsResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
