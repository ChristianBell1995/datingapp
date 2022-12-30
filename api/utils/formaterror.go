package utils

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	if strings.Contains(err, "Duplicate") {
		return errors.New("Record already exists")
	}

	if strings.Contains(err, "not found") {
		return errors.New("Record not found")
	}

	return errors.New("Internal Server Error")
}
