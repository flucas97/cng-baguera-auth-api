package auth_service

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/account"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

var (
	AuthService authServiceInterface = &authService{}
)

type authServiceInterface interface {
	New(account.Login) (string, *error_factory.RestErr)
}

type authService struct{}

func (au *authService) New(account account.Login) (string, *error_factory.RestErr) {
	auth := auth.New(account.Name)
	token, err := auth.GenerateToken()
	if err != nil {
		return "", err
	}
	return token, nil
}
