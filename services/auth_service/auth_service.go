package auth_service

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

var (
	AuthService authServiceInterface = &authService{}
)

type authServiceInterface interface {
	Authorize(string, context.Context) (string, *error_factory.RestErr)
	Validate(string, context.Context) (bool, *error_factory.RestErr)
}

type authService struct{}

func (au *authService) Authorize(nickName string, ctx context.Context) (string, *error_factory.RestErr) {
	token := auth.New(nickName)

	jwt, err := token.GenerateJWT()
	if err != nil {
		return "", err
	}

	if err := token.Authorize(ctx); err != nil {
		return "", err
	}

	return jwt, nil
}

func (au *authService) Validate(jwt string, ctx context.Context) (bool, *error_factory.RestErr) {
	nickName, err := auth.GetNickNameFromJWT(jwt)
	if err != nil {
		return false, err
	}

	ok, err := auth.Validate(ctx, nickName, jwt)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, error_factory.NewNotFoundError("jwt not valid")
	}

	return true, nil
}
