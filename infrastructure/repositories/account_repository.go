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
	repository.db.Create(account)
	return !repository.db.NewRecord(account)
}
