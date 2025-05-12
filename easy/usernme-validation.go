package main

import (
	"regexp"
)

func CodelandUsernameValidation(str string) string {
	f := "false"
	t := "true"

	pattern := `^[A-Za-z][A-Za-z0-9_]{2,23}[A-Za-z0-9]$`
	matched, _ := regexp.MatchString(pattern, str)

	if matched {
		return t
	}

	// code goes here
	return f

}
