package base

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Функция подключается к базе данных, используя драйвер PostgreSQL
func ConnectToDb() (*sql.DB, error) {
	// Имя хоста и пользователя базы данных. Используется GetEnv, так как в Docker-контейнере имена переменных могут отличаться
	databaseHost := GetEnv("DATABASE_HOST", "localhost")
	databaseUser := GetEnv("DATABASE_USER", "student")
	// Строка подключения
	connection := fmt.Sprintf("postgresql://%s:qwerty@%s:5432/proxy_db?sslmode=disable", databaseUser, databaseHost)

	var db *sql.DB
	// Проверка на доступность подключения к бд
	// В самый первый запуск база данных инициализирует себя, поэтому подключение к ней приведет ошибку, так как она еще не готова
	// Этот цикл дает пять секунд на подключение к бд, иначе выдает ошибку и заканчивает работу программы
	connectionAttempt := 0
	attemptsLimit := 5
	for ; connectionAttempt < attemptsLimit; connectionAttempt++ {
		var err error
		db, err = sql.Open("postgres", connection)
		if err != nil {
			LogError(err)
			os.Exit(-1)
		}
		if err := db.Ping(); err != nil {
			LogError(errors.New(
				"database is not responding; attempts: " + strconv.Itoa(connectionAttempt+1) + "/" + strconv.Itoa(attemptsLimit)))
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	if connectionAttempt == 5 {
		return nil, errors.New("failed to connect to database")
	}
	return db, nil
}
