package utils

import "github.com/joyzem/documents/services/base"

const (
	ORGANIZATIONS_ENV = "ORGANIZATION_BACKEND_ADDRESS"
	ACCOUNTS_ENV      = "ACCOUNTS_BACKEND_ADDRESS"
)

func GetOrganizationsAddress() string {
	address := base.GetEnv(ORGANIZATIONS_ENV, "http://localhost:7072")
	return address
}

func GetAccountsAddress() string {
	address := base.GetEnv(ACCOUNTS_ENV, "http://localhost:7073")
	return address
}
