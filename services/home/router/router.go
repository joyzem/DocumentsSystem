package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	productAddress := base.GetEnv("PRODUCT_ADDRESS", "http://localhost:8081")
	handleWithProxy(router, "/product", productAddress)

	organizationAddress := base.GetEnv("ORGANIZATION_ADDRESS", "http://localhost:8082")
	handleWithProxy(router, "/organization", organizationAddress)

	accountAddress := base.GetEnv("ACCOUNT_ADDRESS", "http://localhost:8083")
	handleWithProxy(router, "/account", accountAddress)

	employeeAddress := base.GetEnv("EMPLOYEE_ADDRESS", "http://localhost:8084")
	handleWithProxy(router, "/employee", employeeAddress)

	customerAddress := base.GetEnv("CUSTOMER_ADDRESS", "http://localhost:8085")
	handleWithProxy(router, "/customer", customerAddress)

	proxyAddress := base.GetEnv("PROXY_ADDRESS", "http://localhost:8086")
	handleWithProxy(router, "/documents", proxyAddress)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))

	return router
}

// Redirect from given path to the same path but on the another address
func handleWithProxy(router *mux.Router, path string, address string) {
	router.PathPrefix(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(address)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})
}
