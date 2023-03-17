package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/product/frontend/router"
)

	func main() {
		// получение маршрутизатора
		handler := router.GetRouter()
		// запуск сервера
		fmt.Println("Listening on 8081...")
		if err := http.ListenAndServe(":8081", handler); err != nil {
			base.LogError(err)
			os.Exit(-1)
		}
	}
