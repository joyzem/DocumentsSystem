package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/organization/backend/repo"
	"github.com/joyzem/documents/services/organization/domain"
)

type repository struct {
	db *sql.DB
}

func NewOrganizationRepo(db *sql.DB) repo.OrganizationRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateOrganization(org domain.Organization) (*domain.Organization, error) {
	sql := `
		INSERT INTO organizations (name, address, account_id, chief, financial_chief)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(org.Name, org.Address, org.AccountId, org.Chief, org.FinancialChief).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.OrganizationById(insertedId)
}

func (r *repository) GetOrganizations() ([]domain.Organization, error) {
	sql := `
		SELECT * FROM organizations
	`

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	organizations := []domain.Organization{}

	for rows.Next() {
		organization := domain.Organization{}
		rows.Scan(
			&organization.Id,
			&organization.Name,
			&organization.Address,
			&organization.AccountId,
			&organization.Chief,
			&organization.FinancialChief,
		)
		organizations = append(organizations, organization)
	}

	return organizations, nil
}

func (r *repository) UpdateOrganization(org domain.Organization) (*domain.Organization, error) {
	sql := `
		UPDATE organizations SET name = $1, address = $2, account_id = $3, chief = $4, financial_chief = $5 WHERE id = $6
	`
	rows, err := r.db.Query(sql, org.Name, org.Address, org.AccountId, org.Chief, org.FinancialChief, org.Id)
	if err != nil {
		return nil, err
	}
	rows.Close()

	return r.OrganizationById(org.Id)
}

func (r *repository) DeleteOrganization(id int) error {
	sql := `
		DELETE FROM organizations WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *repository) OrganizationById(id int) (*domain.Organization, error) {

	sql := `
		SELECT * FROM organizations WHERE id = $1
	`

	org := domain.Organization{}

	if err := r.db.QueryRow(sql, id).Scan(
		&org.Id,
		&org.Name,
		&org.Address,
		&org.AccountId,
		&org.Chief,
		&org.FinancialChief); err != nil {

		return nil, err
	}

	return &org, nil
}
