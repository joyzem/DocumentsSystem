package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/product/backend/service"
	"github.com/joyzem/documents/services/product/dto"
)

// Структура которая хранит в себе все эндпоинты, используемые в микросервисе.
type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
	ProductById   endpoint.Endpoint
	CreateUnit    endpoint.Endpoint
	GetUnits      endpoint.Endpoint
	UpdateUnit    endpoint.Endpoint
	DeleteUnit    endpoint.Endpoint
	UnitById      endpoint.Endpoint
}

// Создает и возвращает набор эндпоинтов, используя сервис `s`.
// Для создания каждого эндпоинта используется соответствующая
// функция (например, `makeCreateProductEndpoint` для создания
// эндпоинта `CreateProduct`)
func MakeEndpoints(s service.ProductService) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		GetProducts:   makeGetProductsEndpoint(s),
		UpdateProduct: makeUpdateProductEndpoint(s),
		DeleteProduct: makeDeleteProductEndpoint(s),
		ProductById:   makeProductByIdEndpoint(s),
		CreateUnit:    makeCreateUnitEndpoint(s),
		GetUnits:      makeGetUnitsEndpoint(s),
		UpdateUnit:    makeUpdateUnitEndpoint(s),
		DeleteUnit:    makeDeleteUnitEndpoint(s),
		UnitById:      makeUnitByIdEndpoint(s),
	}
}

// Возвращает endpoint, который создает новый продукт с помощью сервиса `s`
func makeCreateProductEndpoint(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Приводим запрос к типу `dto.CreateProductRequest`
		req := request.(dto.CreateProductRequest)
		// Создаем новый продукт с помощью сервиса
		product, err := s.CreateProduct(req.Name, req.Price, req.UnitId)
		// Возвращаем ответ в виде `dto.CreateProductResponse` с полем `Product`
		return dto.CreateProductResponse{Product: product}, err
	}
}

	// Возвращает endpoint, который получает список всех продуктов с помощью сервиса `s`
	func makeGetProductsEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Получаем список всех продуктов с помощью сервиса
			products, err := s.GetProducts()
			// Возвращаем ответ в виде `dto.GetProductsResponse` с полем `Products`
			return dto.GetProductsResponse{Products: products}, err
		}
	}

	// Возвращает endpoint, который обновляет продукт с помощью сервиса `s`
	func makeUpdateProductEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Приведение к `dto.UpdateProductRequest`
			req := request.(dto.UpdateProductRequest)
			// Обновление информации о товаре
			product, err := s.UpdateProduct(req.Product)
			// Возвращаем ответ в виде `dto.UpdateProductResponse` с полем `Product`
			return dto.UpdateProductResponse{Product: product}, err
		}
	}

	// Возвращает endpoint, который удаляет товар с помощью сервиса `s`
	func makeDeleteProductEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Приводим запрос к типу `dto.DeleteProductRequest`
			req := request.(dto.DeleteProductRequest)
			// Удаляем товар с помощью сервиса
			err := s.DeleteProduct(req.Id)
			// Возвращаем ответ в виде `dto.DeleteProductResponse`
			return dto.DeleteProductResponse{}, err
		}
	}

	// Возвращает endpoint, который возвращает товар по идентификатору
	func makeProductByIdEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(dto.ProductByIdRequest)
			product, err := s.ProductById(req.Id)
			return dto.ProductByIdResponse{Product: product}, err
		}
	}

	// Возвращает endpoint, который удаляет единицу измерения с помощью сервиса `s`
	func makeDeleteUnitEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Приводим запрос к типу `dto.DeleteUnitRequest`
			req := request.(dto.DeleteUnitRequest)
			// Удаляем единицу измерения с помощью сервиса
			err := s.DeleteUnit(req.Id)
			// Возвращаем ответ в виде `dto.DeleteUnitResponse`
			return dto.DeleteUnitResponse{}, err
		}
	}

	// Возвращает endpoint, который обновялет единицу измерения с помощью сервиса `s`
	func makeUpdateUnitEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Приведение к типу
			req := request.(dto.UpdateUnitRequest)
			// Обновление
			unit, err := s.UpdateUnit(req.Unit)
			// Возврат ответа
			return dto.UpdateUnitResponse{Unit: unit}, err
		}
	}

	// Возвращает endpoint, который получает все единицы измерения.
	func makeGetUnitsEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Вызывается функция GetUnits из сервиса и получается список единиц измерения.
			units, err := s.GetUnits()
			// Возвращается ответ в формате GetUnitsResponse со списком единиц измерения и возможной ошибкой.
			return dto.GetUnitsResponse{
				Units: units,
			}, err
		}
	}

	// Возвращает endpoint, который создает новую единицу измерения.
	func makeCreateUnitEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// Преобразуется запрос в тип CreateUnitRequest.
			req := request.(dto.CreateUnitRequest)
			// Вызывается функция CreateUnit из сервиса, передавая параметр единицы измерения.
			unit, err := s.CreateUnit(req.Unit)
			// Возвращается ответ в форме `dto.CreateUnitResponse`
			return dto.CreateUnitResponse{Unit: unit}, err
		}
	}

	// Возвращает endpoint, который возвращает единицу измерения по идентификатору
	func makeUnitByIdEndpoint(s service.ProductService) endpoint.Endpoint {
		return func(_ context.Context, request interface{}) (interface{}, error) {
			req := request.(dto.UnitByIdRequest)
			unit, err := s.UnitById(req.Id)
			return dto.UnitByIdResponse{Unit: unit}, err
		}
	}
