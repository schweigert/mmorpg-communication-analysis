package main

import (
	"github.com/krakenlab/ternary"
	"github.com/schweigert/mmorpg-communication-analysis/clients/client"
	"github.com/schweigert/mmorpg-communication-analysis/clients/wwebproxy"
)

func init() {
	client.FillPersonalData()
}

func main() {
	createAccount()
}

func createAccount() {
	client.FillPersonalData()
	form := client.CredentialForm()

	var err error

	client.Account, err = wwebproxy.CreateAccount(form)
	ternary.Func(err == nil, func() {}, createAccount)()
}
