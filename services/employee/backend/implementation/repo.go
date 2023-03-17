package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/employee/backend/repo"
	"github.com/joyzem/documents/services/employee/domain"
)

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) repo.EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) CreateEmployee(empl domain.Employee) (*domain.Employee, error) {
	sql := `
	INSERT INTO employees (
		first_name,
		last_name,
		middle_name,
		post,
		passport_series,
		passport_number,
		passport_issued_by,
		passport_date_of_issue)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(
		empl.FirstName,
		empl.LastName,
		empl.MiddleName,
		empl.Post,
		empl.PassportSeries,
		empl.PassportNumber,
		empl.PassportIssuedBy,
		empl.PassportDateOfIssue,
	).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.EmployeeById(insertedId)
}

func (r *employeeRepo) GetEmployees() ([]domain.Employee, error) {
	sql := `
		SELECT * FROM employees
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []domain.Employee{}

	for rows.Next() {
		employee := domain.Employee{}
		rows.Scan(
			&employee.Id,
			&employee.FirstName,
			&employee.LastName,
			&employee.MiddleName,
			&employee.Post,
			&employee.PassportSeries,
			&employee.PassportNumber,
			&employee.PassportIssuedBy,
			&employee.PassportDateOfIssue,
		)
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *employeeRepo) EmployeeById(id int) (*domain.Employee, error) {
	sql := `
		SELECT * FROM employees WHERE id = $1
	`
	employee := domain.Employee{}
	if err := r.db.QueryRow(sql, id).Scan(
		&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.MiddleName,
		&employee.Post,
		&employee.PassportSeries,
		&employee.PassportNumber,
		&employee.PassportIssuedBy,
		&employee.PassportDateOfIssue,
	); err != nil {
		return nil, err
	}
	parsedTime, err := base.ParseTime(employee.PassportDateOfIssue)
	employee.PassportDateOfIssue = parsedTime
	return &employee, err
}

func (r *employeeRepo) DeleteEmployee(id int) error {
	sql := `
		DELETE FROM employees WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *employeeRepo) UpdateEmployee(empl domain.Employee) (*domain.Employee, error) {
	sql := `
		UPDATE employees SET 
			first_name = $1,
			last_name = $2,
			middle_name = $3,
			post = $4,
			passport_series = $5,
			passport_number = $6,
			passport_issued_by = $7,
			passport_date_of_issue = $8
		WHERE id = $9
	`
	rows, err := r.db.Query(
		sql,
		empl.FirstName,
		empl.LastName,
		empl.MiddleName,
		empl.Post,
		empl.PassportSeries,
		empl.PassportNumber,
		empl.PassportIssuedBy,
		empl.PassportDateOfIssue,
		empl.Id,
	)
	if err != nil {
		return nil, err
	}
	rows.Close()
	return r.EmployeeById(empl.Id)
}
