package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/product/backend/implementation"
	"github.com/joyzem/documents/services/product/backend/transport"
	httptransport "github.com/joyzem/documents/services/product/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	// Подключение к БД
	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	defer db.Close()

	// Репозиторий товаров
	productRepo := implementation.NewProductRepo(db)
	// Репозиторий единиц измерения
	unitRepo := implementation.NewUnitRepository(db)
	// Создание сервиса
	svc := implementation.NewService(productRepo, unitRepo)

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	// Создание маршрутизатора
	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7071...")
	if err := http.ListenAndServe(":7071", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

}
