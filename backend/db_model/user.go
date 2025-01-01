package db_model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (u User) Validate() error {
	err := validation.ValidateStruct(&u,
		validation.Field(&u.ID,
			validation.Required,
		),
		validation.Field(&u.Name,
			validation.Required,
			validation.Length(1, 64),
		),
		validation.Field(&u.Email,
			validation.Required,
			validation.Length(1, 80),
			is.Email,
		),
	)
	if err != nil {
		return err
	}
	return nil
}
