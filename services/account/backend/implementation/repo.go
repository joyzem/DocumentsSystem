package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/account/backend/repo"
	"github.com/joyzem/documents/services/account/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) repo.AccountRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateAccount(acc domain.Account) (*domain.Account, error) {
	sql := `
		INSERT INTO accounts (account, bank_name, bank_identity_number)
		VALUES ($1, $2, $3) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(acc.Account, acc.BankName, acc.BankIdentityNumber).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.AccountById(insertedId)
}

func (r *repository) GetAccounts() ([]domain.Account, error) {
	sql := `
		SELECT * FROM accounts ORDER BY bank_name ASC
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return []domain.Account{}, err
	}
	defer rows.Close()

	accounts := []domain.Account{}

	for rows.Next() {
		account := domain.Account{}
		rows.Scan(&account.Id, &account.Account, &account.BankName, &account.BankIdentityNumber)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *repository) UpdateAccount(acc domain.Account) (*domain.Account, error) {
	sql := `
		UPDATE accounts 
			SET account = $1, bank_name = $2, bank_identity_number = $3
		WHERE id = $4
	`
	rows, err := r.db.Query(sql, acc.Account, acc.BankName, acc.BankIdentityNumber, acc.Id)
	if err != nil {
		return nil, err
	}

	rows.Close()

	return r.AccountById(acc.Id)
}

func (r *repository) DeleteAccount(id int) error {
	sql := `
		DELETE FROM accounts WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *repository) AccountById(id int) (*domain.Account, error) {
	sql := `
		SELECT * FROM accounts WHERE id = $1
	`
	account := domain.Account{}
	if err := r.db.QueryRow(sql, id).Scan(
		&account.Id,
		&account.Account,
		&account.BankName,
		&account.BankIdentityNumber,
	); err != nil {
		return nil, err
	}
	return &account, nil
}
