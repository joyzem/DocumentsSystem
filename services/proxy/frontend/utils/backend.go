	package utils

	import "github.com/joyzem/documents/services/base"

	const (
		PROXIES_ENV       = "PROXY_BACKEND_ADDRESS"
		EMPLOYEES_ENV     = "EMPLOYEE_BACKEND_ADDRESS"
		PRODUCTS_ENV      = "PRODUCT_BACKEND_ADDRESS"
		ORGANIZATIONS_ENV = "ORGANIZATION_BACKEND_ADDRESS"
		CUSTOMERS_ENV     = "CUSTOMER_BACKEND_ADDRESS"
		ACCOUNTS_ENV      = "ACCOUNT_BACKEND_ADDRESS"
	)

	func GetProxiesAddress() string {
		address := base.GetEnv(PROXIES_ENV, "http://localhost:7076")
		return address
	}

	func GetEmployeesAddress() string {
		address := base.GetEnv(EMPLOYEES_ENV, "http://localhost:7074")
		return address
	}

	func GetProductsAddress() string {
		address := base.GetEnv(PRODUCTS_ENV, "http://localhost:7071")
		return address
	}

	func GetOrganizationsAddress() string {
		address := base.GetEnv(ORGANIZATIONS_ENV, "http://localhost:7072")
		return address
	}

	func GetCustomersAddress() string {
		address := base.GetEnv(CUSTOMERS_ENV, "http://localhost:7075")
		return address
	}

	func GetAccountsAddress() string {
		address := base.GetEnv(ACCOUNTS_ENV, "http://localhost:7073")
		return address
	}
