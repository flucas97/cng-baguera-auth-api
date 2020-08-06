package auth_service

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/account"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
)

var (
	AuthService authServiceInterface = &authService{}
)

type authServiceInterface interface {
	New(account.Login)
}

type authService struct{}

func (au *authService) New(account account.Login) {
	token, err := auth.New(account.Name)
	if err != nil {

	}
	_ = token
}
