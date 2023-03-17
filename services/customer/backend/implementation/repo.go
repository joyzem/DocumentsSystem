package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/customer/backend/repo"
	"github.com/joyzem/documents/services/customer/domain"

	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) repo.CustomerRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCustomer(customer domain.Customer) (*domain.Customer, error) {
	sql := `
		INSERT INTO customers (name) VALUES ($1) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(customer.Name).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.CustomerById(insertedId)
}

func (r *repository) GetCustomers() ([]domain.Customer, error) {
	sql := `
		SELECT * FROM customers
	`

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []domain.Customer{}

	for rows.Next() {
		customer := domain.Customer{}
		rows.Scan(
			&customer.Id,
			&customer.Name,
		)
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *repository) UpdateCustomer(customer domain.Customer) (*domain.Customer, error) {
	sql := `
		UPDATE customers SET name = $1 WHERE id = $2
	`
	rows, err := r.db.Query(sql, customer.Name, customer.Id)
	if err != nil {
		return nil, err
	}

	rows.Close()

	return r.CustomerById(customer.Id)
}

func (r *repository) DeleteCustomer(id int) error {
	sql := `
		DELETE FROM customers WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *repository) CustomerById(id int) (*domain.Customer, error) {
	sql := `
		SELECT * FROM customers WHERE id = $1
	`
	customer := domain.Customer{}
	if err := r.db.QueryRow(sql, id).Scan(&customer.Id, &customer.Name); err != nil {
		return nil, err
	}
	return &customer, nil
}
