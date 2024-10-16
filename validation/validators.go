package validation

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidFloat = errors.New("please enter a valid number")
)

func IsFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return ErrInvalidFloat
	}
	
	return nil
}