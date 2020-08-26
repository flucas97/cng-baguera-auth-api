package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	uuid "github.com/satori/go.uuid"
)

// Token is the model for autheticate user, it has all info about him, and used as interface for JWT cast
type Token struct {
	Name                 string `json:"name"`
	UUID                 string `json:"_uuid"`
	Jwt                  string `json:"jwt"`
	AccountID            string `json:"account_id"`
	CannabisRepositoryID string `json:"cannabis_repository_id"`
}

// New creates a new Token for a given user
func New(name string, accountID string, cannabisRepositoryID string) *Token {
	au := &Token{}
	au.UUID = uuid.NewV4().String()
	au.Name = name
	au.AccountID = accountID
	au.CannabisRepositoryID = cannabisRepositoryID
	return au
}

// GenerateJWT creates a new jwt using the Token information, returns an string of jwt
func (au *Token) GenerateJWT() (string, *error_factory.RestErr) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nick_name"] = au.Name
	claims["uuid"] = au.UUID
	claims["account_id"] = au.AccountID
	claims["cannabis_repository_id"] = au.CannabisRepositoryID

	jwtPure := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := jwtPure.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", error_factory.NewInternalServerError("error generating token")
	}
	au.Jwt = jwt

	return jwt, nil
}

// GetValueFromJwtKey return an given value from an asked key such as: nick_name, uuid, account_id or cannabis_repository_id
func GetValueFromJwtKey(tokens string, key string) (string, *error_factory.RestErr) {
	jwts, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		logger.Error("error parsing jwt", nil)
		return "", error_factory.NewInternalServerError("error getting token informations, try again")
	}

	claims := jwts.Claims.(jwt.MapClaims)
	value := claims[key]

	return (value).(string), nil
}
