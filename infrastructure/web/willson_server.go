package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmorpg-communication-analysis/edge/models"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/forms"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/repositories"
)

// WillsonServer implements wwebservice requests
type WillsonServer struct {
	*Server
	AccountRepository *repositories.AccountRepository
}

// Setup the WillsonServer
func (server *WillsonServer) Setup() {
	server.AccountRepository = repositories.NewAccountRepository()

	server.engine.PUT(AccountRoute, server.PutAccountHandler)
}

// PutAccountHandler create a new account from this web request
func (server *WillsonServer) PutAccountHandler(c *gin.Context) {
	form := &forms.CredentialForm{}

	if c.BindJSON(form) == nil {
		account := models.NewAccount(form.Username, form.Password)
		account.ObfuscatePassword()

		c.JSON(
			ternary.Int(server.AccountRepository.Create(account), http.StatusOK, http.StatusBadRequest),
			account,
		)
	} else {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "JSON Parse error",
			},
		)
	}
}
