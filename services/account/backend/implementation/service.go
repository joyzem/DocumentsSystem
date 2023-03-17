package implementation

import (
	"github.com/joyzem/documents/services/account/backend/repo"
	svc "github.com/joyzem/documents/services/account/backend/service"
	"github.com/joyzem/documents/services/account/domain"
	"github.com/joyzem/documents/services/base"
)

type service struct {
	accountRepo repo.AccountRepo
}

func NewService(accountRepo repo.AccountRepo) svc.Service {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) CreateAccount(account string, bankName string, bankIdentityCode string) (*domain.Account, error) {
	acc := domain.Account{Account: account, BankName: bankName, BankIdentityNumber: bankIdentityCode}
	createdAccount, err := s.accountRepo.CreateAccount(acc)
	base.LogError(err)
	return createdAccount, err
}

func (s *service) GetAccounts() ([]domain.Account, error) {
	accounts, err := s.accountRepo.GetAccounts()
	base.LogError(err)
	return accounts, err
}

func (s *service) UpdateAccount(account domain.Account) (*domain.Account, error) {
	updatedAccount, err := s.accountRepo.UpdateAccount(account)
	base.LogError(err)
	return updatedAccount, err
}

func (s *service) DeleteAccount(id int) error {
	err := s.accountRepo.DeleteAccount(id)
	base.LogError(err)
	return err
}

func (s *service) AccountById(id int) (*domain.Account, error) {
	acc, err := s.accountRepo.AccountById(id)
	base.LogError(err)
	return acc, err
}
