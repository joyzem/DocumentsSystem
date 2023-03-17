package repo

import (
	"github.com/joyzem/documents/services/product/domain"
)

// UnitRepo представляет интерфейс для репозитория единиц измерения.
// Он определяет четыре метода, которые можно выполнять с единицами измерения.
// - CreateUnit - создает новую единицу измерения.
// - GetUnits - получает список всех единиц измерения.
// - UpdateUnit - обновляет единицу измерения.
// - DeleteUnit - удаляет единицу измерения.
// - UnitById - получает единицу измерения по её идентификатору.
type UnitRepo interface {
	CreateUnit(domain.Unit) (*domain.Unit, error)
	GetUnits() ([]domain.Unit, error)
	UpdateUnit(domain.Unit) (*domain.Unit, error)
	DeleteUnit(int) error
	UnitById(int) (*domain.Unit, error)
}
