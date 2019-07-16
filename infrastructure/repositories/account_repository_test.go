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

func TestAccountRepositorySuite(t *testing.T) {
	gspec.Run(t, new(AccountRepositorySuite))
}
