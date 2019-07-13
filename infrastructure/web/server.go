package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/mmorpg-communication-analysis/infrastructure/env"
)

// Server create a default structure for all webservices
type Server struct {
	engine *gin.Engine
}

// Setup root route in this server
func (server *Server) Setup() {
	server.engine.GET(RootRoute, server.Get)
}

// Get root
func (server *Server) Get(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"service": env.ServiceName(),
		},
	)
}
