package base

import "fmt"

func LogError(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
