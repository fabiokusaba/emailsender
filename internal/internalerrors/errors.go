package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("internal server error")

func ProcessErrorToReturn(err error) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return ErrInternal
	}
	return nil
}
