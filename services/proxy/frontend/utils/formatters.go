package utils

import (
	"fmt"
	"strings"
)

func Fullname(lastname, firstname, middlename string) string {
	return fmt.Sprintf(
		"%s %s %s",
		lastname,
		firstname,
		middlename,
	)
}

func NameInitials(fullname string) string {
	splittedName := strings.Split(fullname, " ")
	if len(splittedName) >= 3 {
		lastName := splittedName[0]
		firstNameSlice := []rune(splittedName[1])
		firstName := string(firstNameSlice[0])
		middleNameSlice := []rune(splittedName[2])
		middleName := string(middleNameSlice[0])
		return fmt.Sprintf("%s %s.%s.", lastName, firstName, middleName)
	} else {
		return fullname
	}
}
