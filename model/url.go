package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)
type URL struct {
	gorm.Model
	Address string
	ShortLink string 
}

func (u URL)Validate()error{
	if err := validation.ValidateStruct(
		validation.Field(&u.Address, validation.Required,is.URL),
		validation.Field(&u.ShortLink, validation.Required,validation.Length(1,10)),
	);err!=nil {
		return err;
	}

	return nil;
}