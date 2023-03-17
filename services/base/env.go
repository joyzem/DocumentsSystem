	package base

	import (
		"os"
	)

	// Функция получения переменной из окружения с возможностью возврата значения по умолчанию
	// Требуется для легкого переключения между запуском сервиса через компьютер пользователя или Docker-контейнер
	func GetEnv(key string, defaultValue string) string {
		if value := os.Getenv(key); value == "" {
			return defaultValue
		} else {
			return value
		}
	}
