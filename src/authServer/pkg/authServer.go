package pkg

import (
	"github.com/PenguinCats/MunchkinAuth/src/authServer/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type authServer struct {
	authenticationServer *authentication.AuthenticationServer
	router               *gin.Engine
}

func NewAuthServer() *authServer {
	router := gin.Default()
	authenticateGroup := router.Group("/authenticate")

	authenticationServer := authentication.NewAuthenticationServer()
	authenticationServer.RegisterRoute(authenticateGroup)

	return &authServer{
		authenticationServer: authenticationServer,
		router:               router,
	}
}

func (a *authServer) Start() {
	go func() {
		a.router.Run(":41500")
	}()

	<-make(chan struct{})
}
