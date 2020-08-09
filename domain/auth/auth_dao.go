package auth

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/db/auth/redis_db"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
)

func (token *Token) Authorize(ctx context.Context) *error_factory.RestErr {
	err := redis_db.Client.Set(ctx, token.Name, token.JwtToken, 0).Err()
	if err != nil {
		return error_factory.NewInternalServerError(err.Error())
	}

	return nil
}
