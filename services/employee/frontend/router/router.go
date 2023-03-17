package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/employee/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()
	
	router.HandleFunc("/employee/employees", handlers.EmployeesHandler)

	router.HandleFunc("/employee/employees/create", handlers.CreateEmployeeGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/employee/employees/create", handlers.CreateEmployeePostHandler).Methods(http.MethodPost)

	router.HandleFunc("/employee/employees/update/{id}", handlers.UpdateEmployeeGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/employee/employees/update", handlers.UpdateEmployeePostHandler).Methods(http.MethodPost)

	router.HandleFunc("/employee/employees/detail/{id}", handlers.EmployeeDetailHandler)

	router.HandleFunc("/employee/employees/delete", handlers.DeleteEmployeeHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/employee/static/").Handler(http.StripPrefix("/employee/static/", http.FileServer(http.Dir("../static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})

	handler := c.Handler(router)
	return handler
}
