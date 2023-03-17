package dto

import "github.com/joyzem/documents/services/product/domain"

// Описание DTO (data transfer objects)
// Структуры с тэгом `json:"field"`
// "omitempty" в конце тэга означает, что поле может отсутствовать
// Если запрос выполнен успешно, то придет структура без поля "error",
// иначе — только с полем "error"
type (
	CreateProductRequest struct {
		Name   string `json:"name"`
		Price  int    `json:"price"`
		UnitId int    `json:"unit_id"`
	}
	CreateProductResponse struct {
		Product *domain.Product `json:"product,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	GetProductsRequest struct {
	}
	GetProductsResponse struct {
		Products []domain.Product `json:"products,omitempty"`
		Err      string           `json:"error,omitempty"`
	}
	UpdateProductRequest struct {
		Product domain.Product `json:"product"`
	}
	UpdateProductResponse struct {
		Product *domain.Product `json:"product,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	DeleteProductRequest struct {
		Id int `json:"id"`
	}
	DeleteProductResponse struct {
		Err string `json:"error,omitempty"`
	}
	ProductByIdRequest struct {
		Id int `json:"id"`
	}
	ProductByIdResponse struct {
		Product *domain.Product `json:"product,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
)
