package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "Phone") {
		return errors.New("Phone already registered")
	}

	if strings.Contains(err, "password") {
		return errors.New("Invalid phone or password")
	}

	return errors.New(err)
}
