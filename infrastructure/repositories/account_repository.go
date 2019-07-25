package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmorpg-communication-analysis/edge/models"
)

// AccountRepository over Postgres
type AccountRepository struct {
	db *gorm.DB
}

// Create an account instance
func (repository *AccountRepository) Create(account *models.Account) bool {
	_, ok := repository.FirstWhere("username = ?", account.Username)
	if !ok {
		repository.db.Create(account)
	}

	return !repository.db.NewRecord(account)
}

// FirstWhere element
func (repository *AccountRepository) FirstWhere(query interface{}, args ...interface{}) (*models.Account, bool) {
	account := models.NewAccount("", "")
	repository.db.Where(query, args...).First(account)

	return account, !repository.db.NewRecord(account)
}
