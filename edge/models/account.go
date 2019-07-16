package models

import (
	"crypto/sha256"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Account model
type Account struct {
	gorm.Model

	Username string
	Password string
}

// ObfuscatePassword use Sum256 with sha256 package to obfuscate this password
func (account *Account) ObfuscatePassword() {
	data := []byte("hello")
	hash := sha256.Sum256(data)
	out := fmt.Sprintf("%x", hash[:])
	account.Password = out
}
