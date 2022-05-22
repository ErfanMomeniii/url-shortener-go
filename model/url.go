package model 

import(
	"gorm.io/gorm"
)
type URL struct {
	gorm.Model
	Address string
	ShortLink string 
}
