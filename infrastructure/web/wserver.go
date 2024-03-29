package web

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakenlab/ternary"

	"github.com/schweigert/mmorpg-communication-analysis/edge/models"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/envbuilder"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/forms"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/repositories"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/web/routes"
)

// WServer create a default structure for all webservices
type WServer struct {
	engine *gin.Engine

	accountRepository *repositories.AccountRepository
}

// Setup root route in this Wserver
func (wserver *WServer) Setup() {
	wserver.engine.GET(routes.RootRoute, wserver.GetRootHandler)
	wserver.engine.PUT(routes.AccountRoute, wserver.PutAccountHandler)
}

// Start this Wserver
func (wserver *WServer) Start() {
	panic(wserver.engine.Run(envbuilder.GinAddr()))
}

// BindForm decode data from request body
func (wserver *WServer) BindForm(c *gin.Context, form interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(form)
}

// GetRootHandler function
func (wserver *WServer) GetRootHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"service": env.ServiceName(),
		},
	)
}

// PutAccountHandler function
func (wserver *WServer) PutAccountHandler(c *gin.Context) {
	form := &forms.CredentialForm{}

	if wserver.BindForm(c, form) == nil && form.Valid() {
		account := models.NewAccount(form.Username, form.Password)
		account.ObfuscatePassword()

		c.JSON(
			ternary.Int(wserver.accountRepository.Create(account), http.StatusOK, http.StatusConflict),
			account,
		)
	} else {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Invalid Request",
				"form":  form,
			},
		)
	}
}

// PutCharacterHandler function
func (wserver *WServer) PutCharacterHandler(c *gin.Context) {
	form := &forms.CredentialForm{}

	if wserver.BindForm(c, form) == nil && form.Valid() {
		account, ok := wserver.accountRepository.FirstWhere("username = ?", form.Username)
		if ok && account.Username == form.Username && account.Password == account.Password {

		}
	}
}
