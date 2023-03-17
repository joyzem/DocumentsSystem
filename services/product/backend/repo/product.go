package repo

import (
	"github.com/joyzem/documents/services/product/domain"
)

// ProductRepo представляет интерфейс репозитория для управления товарами в базе данных.
// Интерфейс содержит следующие методы:
// - CreateProduct: создает новый товар в базе данных.
// - GetProducts: возвращает список всех товаров в базе данных.
// - UpdateProduct: обновляет информацию о товаре в базе данных.
// - DeleteProduct: удаляет товар из базы данных по его идентификатору.
// - ProductById: получает товар по его идентификатору.
type ProductRepo interface {
	CreateProduct(domain.Product) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	DeleteProduct(int) error
	ProductById(int) (*domain.Product, error)
}
