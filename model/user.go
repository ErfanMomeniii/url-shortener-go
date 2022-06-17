package model

import(
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
}

func (u User)Validate()error{
	if err := validation.ValidateStruct(
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Email, validation.Required,is.Email),
	);err!=nil {
		return err;
	}

	return nil;
}