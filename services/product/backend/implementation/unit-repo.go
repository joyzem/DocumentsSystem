package implementation

import (
	"database/sql"
	"errors"
	"sort"

	"github.com/joyzem/documents/services/product/backend/repo"
	"github.com/joyzem/documents/services/product/domain"
)

// Реализация репозитория единиц измерения
type unitRepository struct {
	db *sql.DB
}

// Возвращает репозиторий единиц измерения
func NewUnitRepository(db *sql.DB) repo.UnitRepo {
	return &unitRepository{
		db: db,
	}
}

// CreateUnit добавляет новую единицу измерения в базу данных и возвращает указатель на созданную единицу измерения или ошибку, если она возникает.
func (r *unitRepository) CreateUnit(unit domain.Unit) (*domain.Unit, error) {
	// Запрос на добавление единицы измерения с возвратом id
	sql := `
			INSERT INTO units (name) VALUES ($1) RETURNING id
	`
	// Подготовка запроса
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Выполнение запроса и сканирование результата в insertedId
	var insertedId int
	if err := result.QueryRow(unit.Name).Scan(&insertedId); err != nil {
		return nil, err
	}

	// Вернуть новую единицу измерения
	return r.UnitById(insertedId)
}

// GetUnits возвращает список всех единиц измерения из базы данных или ошибку, если она возникает.
func (r *unitRepository) GetUnits() ([]domain.Unit, error) {
	// Формирование SQL запроса для получения всех единиц измерения из таблицы units
	sql := `SELECT * FROM units`
	// Выполнение запроса
	rows, err := r.db.Query(sql)
	if err != nil {
		// Возврат ошибки, если произошла ошибка при выполнении запроса
		return nil, err
	}
	// Освобождение ресурсов, занятых результатом запроса
	defer rows.Close()

	// Инициализация массива единиц измерения
	units := []domain.Unit{}

	// Обработка каждой строки результата запроса
	for rows.Next() {
		// Инициализация структуры единицы измерения
		unit := domain.Unit{}
		// Сканирование данных из текущей строки результата запроса
		rows.Scan(&unit.Id, &unit.Name)
		// Добавление единицы измерения в массив
		units = append(units, unit)
	}

	// Сортировка по id
	sort.Slice(units, func(i, j int) bool {
		return units[i].Id < units[j].Id
	})

	return units, nil
}

// UpdateUnit обновляет единицу измерения в базе данных.
// unit - структура с новыми значениями единицы измерения.
// Возвращает указатель на обновленную единицу измерения или ошибку, если она возникла.
func (r *unitRepository) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	// SQL-запрос для обновления информации об единице измерения
	sql := `
			UPDATE units SET name = $1 WHERE id = $2
	`
	// Выполнение запроса
	rows, err := r.db.Query(sql, unit.Name, unit.Id)
	if err != nil {
		return nil, err
	}

	// Закрыть соединение, чтобы осуществить следующий запрос
	rows.Close()

	// Вернуть обновленную единицу измерения
	return r.UnitById(unit.Id)
}

// DeleteUnit выполняет удаление единицы измерения из базы данных по её ID.
// ID единицы измерения передается в качестве параметра.
func (r *unitRepository) DeleteUnit(id int) error {
	if id == 0 {
		return errors.New("can't delete default value")
	}
	sql := `
			DELETE FROM units WHERE id = $1		
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *unitRepository) UnitById(id int) (*domain.Unit, error) {
	// SQL-запрос для получения единицы измерения
	sql := `
			SELECT * FROM units WHERE id = $1	
	`

	// Инициализация переменной
	unit := domain.Unit{}
	if err := r.db.QueryRow(sql, id).Scan(&unit.Id, &unit.Name); err != nil {
		return nil, err
	}

	return &unit, nil
}
