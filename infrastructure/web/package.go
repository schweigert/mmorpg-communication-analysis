package web

import "github.com/gin-gonic/gin"

// Routes
const (
	RootRoute = "/"
)

// Routes include *Route consts
var Routes = []string{
	RootRoute,
}

// NewServer constructor
func NewServer() (server *Server) {
	ginEngine := gin.Default()
	server = &Server{engine: ginEngine}
	server.Setup()
	return server
}

// NewWillsonServer constructor
func NewWillsonServer() *WillsonServer {
	server := NewServer()

	return &WillsonServer{Server: server}
}
