package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmorpg-communication-analysis/edge/models"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/envbuilder"

	// Used to connect in postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", envbuilder.PostgresAddr())
	ternary.Func(err != nil, func() { panic(err) }, func() { db.LogMode(true) })()
}

func availableModels() []interface{} {
	list := []interface{}{}

	list = append(list, &models.Account{})

	return list
}

func migrate() {
	db.AutoMigrate(
		availableModels()...,
	)
}

// NewAccountRepository instance a repository using default db connection
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{db: db}
}
