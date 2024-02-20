package authentication

import "github.com/gin-gonic/gin"

type AuthenticationServer struct {
	mapAuthenticater map[AuthenticaterType]IAuthenticater
}

func NewAuthenticationServer() *AuthenticationServer {
	return &AuthenticationServer{
		mapAuthenticater: make(map[AuthenticaterType]IAuthenticater),
	}
}

func (a *AuthenticationServer) RegisterRoute(routerGroup *gin.RouterGroup) {
	for authenticateType, authenticater := range a.mapAuthenticater {
		subPath, err := authenticateType.GetUrlPath()
		if err != nil {
			panic(err)
		}

		routerGroup.POST(subPath, authenticater.Authenticate)
	}
}
