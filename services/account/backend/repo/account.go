package repo

import (
	"github.com/joyzem/documents/services/account/domain"
)

type AccountRepo interface {
	CreateAccount(domain.Account) (*domain.Account, error)
	GetAccounts() ([]domain.Account, error)
	UpdateAccount(domain.Account) (*domain.Account, error)
	DeleteAccount(int) error
	AccountById(int) (*domain.Account, error)
}
