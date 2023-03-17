package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/account/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/account/accounts", handlers.AccountsHandler)

	router.HandleFunc("/account/accounts/create", handlers.CreateAccountGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/account/accounts/create", handlers.CreateAccountPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/account/accounts/update/{id}", handlers.UpdateAccountGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/account/accounts/update", handlers.UpdateAccountPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/account/accounts/delete", handlers.DeleteAccountHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/account/static/").Handler(http.StripPrefix("/account/static/", http.FileServer(http.Dir("../static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	
	handler := c.Handler(router)
	return handler
}
