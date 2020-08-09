package auth_service

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

var (
	AuthService authServiceInterface = &authService{}
)

type authServiceInterface interface {
	Authorize(string) *error_factory.RestErr
}

type authService struct{}

func (au *authService) Authorize(nickName string) *error_factory.RestErr {
	if nickName == "" {
		return error_factory.NewBadRequestError("account not found")
	}

	token := auth.New(nickName)

	err := token.GenerateJWT()
	if err != nil {
		return err
	}

	if err := token.Authorize(); err != nil {

	}

}
