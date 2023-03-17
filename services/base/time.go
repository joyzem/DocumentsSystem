package base

import (
	"time"
)

func ParseTime(source string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, source)
	if err == nil {
		return parsedTime.Format("2006-01-02"), nil
	}
	return "", err
}
