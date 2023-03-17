package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/proxy/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/documents/proxies", handlers.ProxiesHandler)

	router.HandleFunc("/documents/proxies/create", handlers.CreateProxyGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/proxies/create", handlers.CreateProxyPosthandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/proxies/update/{id:[0-9]+}", handlers.UpdateProxyGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/proxies/update", handlers.UpdateProxyPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/proxies/delete", handlers.DeleteProxyHandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/proxies/details/{id:[0-9]+}", handlers.ProxyDetailsHandler).Methods(http.MethodGet)

	router.HandleFunc("/documents/proxies/{id:[0-9]+}/body/create", handlers.CreateProxyBodyGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/proxies/body/create", handlers.CreateProxyBodyPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/documents/proxies/body/delete", handlers.DeleteProxyBodyHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/documents/static/").Handler(http.StripPrefix("/documents/static/", http.FileServer(http.Dir("../../../documents/static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})

	handler := c.Handler(router)
	return handler
}
