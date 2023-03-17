	package http

	import (
		"context"
		"net/http"
		"strconv"

		kithttp "github.com/go-kit/kit/transport/http"
		"github.com/gorilla/mux"
		"github.com/joyzem/documents/services/base"
		"github.com/joyzem/documents/services/product/backend/transport"
		"github.com/joyzem/documents/services/product/dto"
	)

	// NewService создает новый HTTP-сервис, используя указанные входные точки transport.Endpoints и опции kithttp.ServerOption
	func NewService(
		svcEndpoints transport.Endpoints,
		options []kithttp.ServerOption,
	) http.Handler {

		// Создается новый роутер mux
		router := mux.NewRouter()

		// Устанавливается кодировщик ошибок errorEncoder, который будет использоваться для кодирования ответов об ошибках
		errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)

		// Добавляется кодировщик ошибок в список опций
		options = append(options, errorEncoder)

		// Для следующих методов и путей создаются обработчики kithttp.NewServer, которые реализуют интерфейс http.Handler
		// Каждому методу (POST, GET, DELETE, PUT) и пути (/products, /units) создается обработчик функцией kithttp.NewServer,
		// которая принимает эндпоинт, функцию, декодирующую запрос, функцию, кодирующую ответ и опции
		// Одной из опций является кодирование ошибок. Если в процессе работы сервиса произойдет ошибка, то обработчик вернет
		// JSON ответ в виде {"error": "описание ошибки"}

		router.Methods("POST").Path("/products").Handler(
			kithttp.NewServer(
				svcEndpoints.CreateProduct,
				decodeCreateProductRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("GET").Path("/products").Handler(
			kithttp.NewServer(
				svcEndpoints.GetProducts,
				decodeGetProductsRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("PUT").Path("/products").Handler(
			kithttp.NewServer(
				svcEndpoints.UpdateProduct,
				decodeUpdateProductRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("DELETE").Path("/products").Handler(
			kithttp.NewServer(
				svcEndpoints.DeleteProduct,
				decodeDeleteProductRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("GET").Path("/products/{id}").Handler(
			kithttp.NewServer(
				svcEndpoints.ProductById,
				decodeProductByIdRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("POST").Path("/units").Handler(
			kithttp.NewServer(
				svcEndpoints.CreateUnit,
				decodeCreateUnitRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("GET").Path("/units").Handler(
			kithttp.NewServer(
				svcEndpoints.GetUnits,
				decodeGetUnitsRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("PUT").Path("/units").Handler(
			kithttp.NewServer(
				svcEndpoints.UpdateUnit,
				decodeUpdateUnitRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("DELETE").Path("/units").Handler(
			kithttp.NewServer(
				svcEndpoints.DeleteUnit,
				decodeDeleteUnitRequest,
				base.EncodeResponse,
				options...,
			))

		router.Methods("GET").Path("/units/{id}").Handler(
			kithttp.NewServer(
				svcEndpoints.UnitById,
				decodeUnitByIdRequest,
				base.EncodeResponse,
				options...,
			))

		return router
	}

	// Следующие функции декодируют запрос и возвращают требуемую структуру данных для дальнейшей работы в эндпоинте
	// Приходит http.Request r, и он декодируется с помощью функции base.DecodeBody(r, &req), которая
	// возвращает требуемую структуру или ошибку, если она произошла

	func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.CreateProductRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
		// У этого запроса нет тела, и декодировать нечего
		return dto.GetProductsRequest{}, nil
	}

	func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.UpdateProductRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.DeleteProductRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeProductByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		req := dto.ProductByIdRequest{Id: id}
		return req, nil
	}

	func decodeCreateUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.CreateUnitRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeGetUnitsRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.GetUnitsRequest
		return req, nil
	}

	func decodeUpdateUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.UpdateUnitRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeDeleteUnitRequest(_ context.Context, r *http.Request) (interface{}, error) {
		var req dto.DeleteUnitRequest
		err := base.DecodeBody(r, &req)
		return req, err
	}

	func decodeUnitByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		req := dto.UnitByIdRequest{Id: id}
		return req, nil
	}
