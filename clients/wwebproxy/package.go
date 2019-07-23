package wwebproxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/krakenlab/ternary"
	"github.com/parnurzeal/gorequest"
	"github.com/schweigert/mmorpg-communication-analysis/edge/models"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/forms"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/web/routes"
)

// CreateAccount using gorequest
func CreateAccount(credentialForm *forms.CredentialForm) (*models.Account, error) {
	req, body, errs := gorequest.New().Timeout(5 * time.Second).Put(wWebServiceAccountURL()).Send(credentialForm).End()
	panicErrors(errs...)

	if req.StatusCode == http.StatusOK {
		account := models.NewAccount(credentialForm.Username, credentialForm.Password)
		json.Unmarshal([]byte(body), account)
		return account, nil
	}

	return nil, errors.New(req.Status)
}

func wWebServiceAccountURL() string {
	return fmt.Sprintf("%s%s", wWebServiceURL(), routes.AccountRoute)
}

func wWebServiceURL() string {
	return fmt.Sprintf("%s:%s", wWebServiceHost(), wWebServicePort())
}

func wWebServicePort() string {
	def := os.Getenv("W_WEB_SERVICE_PORT")
	return ternary.String(def != "", def, "3011")
}

func wWebServiceHost() string {
	def := os.Getenv("W_WEB_SERVICE_HOST")
	return ternary.String(def != "", def, "localhost")
}

func panicErrors(errs ...error) {
	if errs == nil {
		return
	}
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
