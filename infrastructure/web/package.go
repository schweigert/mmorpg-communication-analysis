package web

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/repositories"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// NewWServer constructor
func NewWServer() (server *WServer) {
	ginEngine := gin.Default()
	server = &WServer{engine: ginEngine, accountRepository: repositories.NewAccountRepository()}
	server.Setup()
	return server
}
