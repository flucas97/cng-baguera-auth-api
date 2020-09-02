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
	Authorize(context.Context, string, string, string) (string, *error_factory.RestErr)
	Validate(context.Context, string) (bool, *error_factory.RestErr)
}

type authService struct{}

func (au *authService) Authorize(ctx context.Context, nickName string, accountID string, cannabisRepositoryID string) (string, *error_factory.RestErr) {
	token := auth.New(nickName, accountID, cannabisRepositoryID)

	jwt, err := token.GenerateJWT()
	if err != nil {
		return "", err
	}

	if err := token.Authorize(ctx); err != nil {
		return "", err
	}

	return jwt, nil
}

func (au *authService) Validate(ctx context.Context, jwt string) (bool, *error_factory.RestErr) {
	nickName, err := auth.GetValueFromJwtKey(jwt, "nick_name")
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
