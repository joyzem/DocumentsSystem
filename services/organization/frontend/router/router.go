package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/organization/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/organization/organizations", handlers.OrganizationsHandler)

	router.HandleFunc("/organization/organizations/create", handlers.CreateOrganizationHandler).Methods(http.MethodGet)
	router.HandleFunc("/organization/organizations/create", handlers.CreateOrganizationPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/organization/organizations/update/{id}", handlers.UpdateOrganizationGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/organization/organizations/update", handlers.UpdateOrganizationPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/organization/organizations/detail/{id}", handlers.OrganizationDetailHandler)

	router.HandleFunc("/organization/organizations/delete", handlers.DeleteOrganizationHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/organization/static/").Handler(http.StripPrefix("/organization/static/", http.FileServer(http.Dir("../static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})

	handler := c.Handler(router)
	return handler
}
