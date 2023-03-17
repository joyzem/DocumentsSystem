package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/customer/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/customer/customers", handlers.CustomersHandler)

	router.HandleFunc("/customer/customers/create", handlers.CreateCustomerGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer/customers/create", handlers.CreateCustomerPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/customer/customers/update/{id}", handlers.UpdateCustomerGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer/customers/update", handlers.UpdateCustomerPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/customer/customers/delete", handlers.DeleteCustomerHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/customer/static/").Handler(http.StripPrefix("/customer/static/", http.FileServer(http.Dir("../static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})

	handler := c.Handler(router)
	return handler
}
