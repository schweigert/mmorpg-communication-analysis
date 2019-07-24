package repositories

import (
	"testing"

	"github.com/schweigert/mmorpg-communication-analysis/edge/models"

	"github.com/krakenlab/gspec"
)

type AccountRepositorySuite struct {
	gspec.Suite
}

func (suite *AccountRepositorySuite) SetupTest() {
	ClearDB()

	suite.Suite.SetupTest()
}

func (suite *AccountRepositorySuite) TestCreate() {
	account := models.NewAccount("username", "password")
	suite.True(NewAccountRepository().Create(account))

	count := 0
	db.Model(&models.Account{}).Count(&count)
	suite.Equal(1, count)
}

func (suite *AccountRepositorySuite) TestFirstWhere() {
	account := models.NewAccount("username", "password")
	suite.True(NewAccountRepository().Create(account))

	otherAccount, ok := NewAccountRepository().FirstWhere("username = ?", account.Username)

	suite.True(ok)
	suite.Equal(account.Password, otherAccount.Password)
}

func TestAccountRepositorySuite(t *testing.T) {
	gspec.Run(t, new(AccountRepositorySuite))
}
