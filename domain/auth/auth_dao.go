package auth

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/db/auth/redis_db"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
)

func (token *Token) Authorize(ctx context.Context) *error_factory.RestErr {
	_, err := redis_db.Client.Set(ctx, token.Name, token.Jwt, 0).Result()
	if err != nil {
		logger.Error("redis set error", err)
		return error_factory.NewInternalServerError(err.Error())
	}

	return nil
}

func Validate(ctx context.Context, nickName string, jwtSent string) (bool, *error_factory.RestErr) {
	validJWT, err := redis_db.Client.Get(ctx, nickName).Result()
	if err != nil {
		return false, error_factory.NewBadRequestError(err.Error())
	}

	if validJWT != jwtSent {
		return false, error_factory.NewNotFoundError("jwt not valid")
	}

	return true, nil
}
