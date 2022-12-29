package utils

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "name") {
		return errors.New("Name Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	if strings.Contains(err, "Duplicate") {
		return errors.New("Record already exists")
	}
	return errors.New("Internal Server Error")
}
