package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/product/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	// создание маршрутизатора
	router := mux.NewRouter()
	// обработка GET запроса. Выдача страницы всех товаров
	router.HandleFunc("/product/products", handlers.ProductsHandler)
	// обработка POST запроса. Удаление товара
	router.HandleFunc("/product/products/delete", handlers.DeleteProductHandler).Methods(http.MethodPost)
	// обработка GET запроса. Страница добавления товара
	router.HandleFunc("/product/products/create", handlers.CreateProductGetHandler).Methods(http.MethodGet)
	// обработка POST запроса. Добавление товара
	router.HandleFunc("/product/products/create", handlers.CreateProductPostHandler).Methods(http.MethodPost)
	// обработка GET запроса. Страница изменения товара
	router.HandleFunc("/product/products/update/{id}", handlers.UpdateProductGetHandler)
	// обработка POST запроса. Изменение товара
	router.HandleFunc("/product/products/update", handlers.UpdateProductPostHandler).Methods(http.MethodPost)

	// обработка GET запроса. Выдача страницы всех единиц измерения
	router.HandleFunc("/product/units", handlers.UnitsHandler)
	// обработка POST запроса. Удаление единицы измерения
	router.HandleFunc("/product/units/delete", handlers.DeleteUnitHandler).Methods(http.MethodPost)
	// обработка GET запроса. Страница добавления единицы измерения
	router.HandleFunc("/product/units/create", handlers.CreateUnitGetHandler).Methods(http.MethodGet)
	// обработка POST запроса. Добавление единицы измерения
	router.HandleFunc("/product/units/create", handlers.CreateUnitPostHandler).Methods(http.MethodPost)
	// обработка GET запроса. Страница изменения единицы измерения
	router.HandleFunc("/product/units/update/{id}", handlers.UpdateUnitGetHandler)
	// обработка POST запроса. Изменение единицы измерения
	router.HandleFunc("/product/units/update", handlers.UpdateUnitPostHandler).Methods(http.MethodPost)

	// обработка обращения к глобальным статическим ресурсам: js, css. Используется в html
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	// обработка обращения к локальным статическим ресурсам
	router.PathPrefix("/product/static/").Handler(http.StripPrefix("/product/static/", http.FileServer(http.Dir("../static"))))

	// механизм, позволяющий отправлять запросы на другой домен (адрес)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})
	// обертка над маршрутизатором
	handler := c.Handler(router)
	return handler
}
