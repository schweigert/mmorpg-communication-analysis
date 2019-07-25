package models

import (
	"github.com/jinzhu/gorm"
)

// Character model
type Character struct {
	gorm.Model

	AccountID int
	Account   Account

	Name string
}
