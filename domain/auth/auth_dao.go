package auth

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/db/auth/redis_db"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

func (token *Token) Authorize(ctx context.Context) *error_factory.RestErr {
	err := redis_db.Client.Set(ctx, token.Name, token.Jwt, 0).Err()
	if err != nil {
		return error_factory.NewInternalServerError(err.Error())
	}

	return nil
}

func (token *Token) Validate(context.Context) (bool, *error_factory.RestErr) {
	/*
		check if is any register of this account in Redis
		then check if the jwt token is the same in Token.Jwt
	*/
	return true, nil
}
