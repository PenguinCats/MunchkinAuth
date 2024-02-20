package authentication

import (
	"errors"

	"github.com/PenguinCats/MunchkinAuth/src/authServer/pkg/common/constants"
	"github.com/gin-gonic/gin"
)

type AuthenticaterType int

const (
	OPENIDMICROSOFT AuthenticaterType = iota
	OPENIDGOOGLE
	OPENIDTencent
	USERNAMEPASSWORD
)

func (a AuthenticaterType) GetUrlPath() (string, error) {
	switch a {
	case OPENIDMICROSOFT:
		return constants.AUTHENTICATORTYPEOPENIDMICROSOFTSUBPATH, nil
	case OPENIDGOOGLE:
		return constants.AUTHENTICATORTYPEOPENIDGOOGLESUBPATH, nil
	case OPENIDTencent:
		return constants.AUTHENTICATORTYPEOPENIDTENCENTSUBPATH, nil
	case USERNAMEPASSWORD:
		return constants.AUTHENTICATORTYPEUSERNAMEPASSWORDSUBPATH, nil
	default:
		return "", errors.New("unknown authenticater type")
	}
}

type IAuthenticater interface {
	Authenticate(*gin.Context)
}
