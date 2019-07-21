package models

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmorpg-communication-analysis/edge/password"
)

// Account model
type Account struct {
	gorm.Model

	Username string
	Password string
}

// ObfuscatePassword use Sum256 with sha256 package to obfuscate this password
func (account *Account) ObfuscatePassword() {
	account.Password = password.Sum(account.Password)
}
