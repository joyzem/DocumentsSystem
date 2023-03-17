package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/backend/repo"
	"github.com/joyzem/documents/services/proxy/domain"
)

type proxyRepo struct {
	db *sql.DB
}

func NewProxyRepo(db *sql.DB) repo.ProxyRepo {
	return &proxyRepo{
		db: db,
	}
}

func (r *proxyRepo) CreateProxyHeader(header domain.ProxyHeader) (*domain.ProxyHeader, error) {
	sql := `
		INSERT INTO proxies 
		(
			organization_id,
			customer_id,
			employee_id,
			date_of_issue,
			is_valid_until
		)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(
		header.OrganizationId,
		header.CustomerId,
		header.EmployeeId,
		header.DateOfIssue,
		header.IsValidUntil,
	).Scan(
		&insertedId,
	); err != nil {
		return nil, err
	}

	header.Id = insertedId
	return &header, nil
}
func (r *proxyRepo) GetProxies() ([]domain.ProxyHeader, error) {
	sql := `
		SELECT * FROM proxies
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	headers := []domain.ProxyHeader{}
	for rows.Next() {
		header := domain.ProxyHeader{}
		if err := rows.Scan(
			&header.Id,
			&header.OrganizationId,
			&header.CustomerId,
			&header.EmployeeId,
			&header.DateOfIssue,
			&header.IsValidUntil,
		); err != nil {
			return nil, err
		}
		header.DateOfIssue, _ = base.ParseTime(header.DateOfIssue)
		header.IsValidUntil, _ = base.ParseTime(header.IsValidUntil)
		headers = append(headers, header)
	}

	return headers, err
}

	func (r *proxyRepo) ProxyById(id int) (*domain.Proxy, error) {
		sql := `
			SELECT * FROM proxies WHERE id = $1
		`
		proxy := domain.Proxy{}
		if err := r.db.QueryRow(sql, id).Scan(
			&proxy.ProxyHeader.Id,
			&proxy.ProxyHeader.OrganizationId,
			&proxy.ProxyHeader.CustomerId,
			&proxy.ProxyHeader.EmployeeId,
			&proxy.ProxyHeader.DateOfIssue,
			&proxy.ProxyHeader.IsValidUntil,
		); err != nil {
			return nil, err
		}

		proxy.ProxyHeader.DateOfIssue, _ = base.ParseTime(proxy.ProxyHeader.DateOfIssue)
		proxy.ProxyHeader.IsValidUntil, _ = base.ParseTime(proxy.ProxyHeader.IsValidUntil)

		getBodySql := `SELECT * FROM proxy_bodies WHERE proxy_id = $1`
		rows, err := r.db.Query(getBodySql, id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			bodyItem := domain.ProxyBodyItem{}
			if err := rows.Scan(
				&bodyItem.Id,
				&bodyItem.ProductId,
				&bodyItem.ProxyId,
				&bodyItem.ProductAmount,
			); err != nil {
				return nil, err
			}
			proxy.ProxyBodyItems = append(proxy.ProxyBodyItems, bodyItem)
		}

		return &proxy, nil
	}

func (r *proxyRepo) UpdateProxyHeader(header domain.ProxyHeader) (*domain.Proxy, error) {
	sql := `
		UPDATE proxies SET 
			organization_id = $1,
			customer_id = $2,
			employee_id = $3,
			date_of_issue = $4,
			is_valid_until = $5
		WHERE id = $6
	`
	rows, err := r.db.Query(
		sql,
		header.OrganizationId,
		header.CustomerId,
		header.EmployeeId,
		header.DateOfIssue,
		header.IsValidUntil,
		header.Id,
	)

	if err != nil {
		return nil, err
	}
	rows.Close()

	return r.ProxyById(header.Id)
}

func (r *proxyRepo) DeleteProxy(id int) error {
	sql := `
		DELETE FROM proxies WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *proxyRepo) CreateProxyBodyItem(item domain.ProxyBodyItem) (*domain.Proxy, error) {
	sql := `
		INSERT INTO proxy_bodies
		(
			product_id, 
			proxy_id, 
			product_amount
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(item.ProductId, item.ProxyId, item.ProductAmount).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.ProxyById(item.ProxyId)
}

func (r *proxyRepo) DeleteProxyBodyItem(id int) error {
	sql := `
		DELETE FROM proxy_bodies WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}
