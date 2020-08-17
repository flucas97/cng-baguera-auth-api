package auth

import (
	"context"
	"strings"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/db/auth/psql_db"
	"github.com/flucas97/cng/cng-baguera-auth-api/db/auth/redis_db"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
)

const (
	queryNewCannabis = ("INSERT INTO auth (account_id, jwt, nickname) VALUES ($1, $2, $3);")
	queryShowDetails = ("SELECT jwt FROM auth WHERE nickname=$1;")
)

func (token *Token) Authorize(ctx context.Context) *error_factory.RestErr {
	_, err := redis_db.Client.Set(ctx, token.Name, token.Jwt, 600000*time.Second).Result()
	if err != nil {
		logger.Error("redis set error", err)
		return error_factory.NewInternalServerError(err.Error())
	}

	_ = psql_db.Client.QueryRow(queryNewCannabis, token.AccountId, token.Jwt, token.Name)

	return nil
}

func Validate(ctx context.Context, nickName string, jwtSent string) (bool, *error_factory.RestErr) {
	validJWT, err := redis_db.Client.Get(ctx, nickName).Result()
	if err != nil {
		err = psql_db.Client.QueryRow(queryShowDetails, nickName).Scan(&validJWT)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				return false, error_factory.NewNotFoundError("account not found")
			}
			logger.Error("error while preparing query", err)
			return false, error_factory.NewInternalServerError("error searching account informations, try again")
		}
	}

	if validJWT != jwtSent {
		return false, error_factory.NewNotFoundError("jwt not valid")
	}

	return true, nil
}
