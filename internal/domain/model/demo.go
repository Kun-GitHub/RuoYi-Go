package model

import (
	"gorm.io/gorm"
)

type Demo struct {
	gorm.Model
	Name  string
	Email string
}
