package implementation

import (
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/product/backend/repo"
	"github.com/joyzem/documents/services/product/backend/service"
	"github.com/joyzem/documents/services/product/domain"
)

// Реализация сервиса
type productService struct {
	productRepo repo.ProductRepo
	unitRepo    repo.UnitRepo
}

// Возвращает реализацию сервиса
func NewService(productRepo repo.ProductRepo, unitRepo repo.UnitRepo) service.ProductService {
	return &productService{
		productRepo: productRepo,
		unitRepo:    unitRepo,
	}
}

// CreateProduct - создает новый товар с переданными параметрами (name, price, unitId)
func (s *productService) CreateProduct(name string, price int, unitId int) (*domain.Product, error) {
	p := domain.Product{Name: name, Price: price, UnitId: unitId}
	createdProduct, err := s.productRepo.CreateProduct(p)
	base.LogError(err)
	return createdProduct, err
}

// GetProducts - возвращает все товары
func (s *productService) GetProducts() ([]domain.Product, error) {
	products, err := s.productRepo.GetProducts()
	base.LogError(err)
	return products, err
}

// UpdateProduct - обновляет информацию о товаре
func (s *productService) UpdateProduct(newProduct domain.Product) (*domain.Product, error) {
	updatedProduct, err := s.productRepo.UpdateProduct(newProduct)
	base.LogError(err)
	return updatedProduct, err
}

// DeleteProduct - удаляет товар по id
func (s *productService) DeleteProduct(id int) error {
	err := s.productRepo.DeleteProduct(id)
	base.LogError(err)
	return err
}

// ProductById - возвращает товар по идентификатору
func (s *productService) ProductById(id int) (*domain.Product, error) {
	product, err := s.productRepo.ProductById(id)
	base.LogError(err)
	return product, err
}

// CreateUnit - добавляет единицу измерения.
func (s *productService) CreateUnit(name string) (*domain.Unit, error) {
	unit := domain.Unit{Name: name}
	createdUnit, err := s.unitRepo.CreateUnit(unit)
	base.LogError(err)
	return createdUnit, err
}

// GetUnits - возвращает список всех единиц измерения.
func (s *productService) GetUnits() ([]domain.Unit, error) {
	units, err := s.unitRepo.GetUnits()
	base.LogError(err)
	return units, err
}

// UpdateUnit - обновляет информацию об единице измерения.
func (s *productService) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	updatedUnit, err := s.unitRepo.UpdateUnit(unit)
	base.LogError(err)
	return updatedUnit, err
}

// DeleteUnit - удаляет единицу измерения.
func (s *productService) DeleteUnit(id int) error {
	err := s.unitRepo.DeleteUnit(id)
	base.LogError(err)
	return err
}

// UnitById - возвращает единицу измерения по идентификатору
func (s *productService) UnitById(id int) (*domain.Unit, error) {
	unit, err := s.unitRepo.UnitById(id)
	base.LogError(err)
	return unit, err
}
