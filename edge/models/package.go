package models

// NewAccount using username and passowrd
func NewAccount(username string, securePassowrd string) *Account {
	return &Account{Username: username, Password: securePassowrd}
}
