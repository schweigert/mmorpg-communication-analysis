package models

import (
	"testing"

	"github.com/krakenlab/gspec"
)

type AccountTest struct {
	gspec.Suite
}

func (suite *AccountTest) TestObfuscatePassword() {
	account := NewAccount("username", "securePass")
	suite.NotPanics(account.ObfuscatePassword)
	suite.NotEqual("securePassword", account.Password)
	suite.Equal("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", account.Password)
}

func TestAccountTest(t *testing.T) {
	gspec.Run(t, new(AccountTest))
}
