package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	uuid "github.com/satori/go.uuid"
)

type Token struct {
	Name string `json:"name"`
	Uuid string `json:"_uuid"`
	Jwt  string `json:"jwt"`
}

func New(name string) *Token {
	au := &Token{}
	au.Uuid = uuid.NewV4().String()
	au.Name = name
	return au
}

func (au *Token) GenerateJWT() (string, *error_factory.RestErr) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nick_name"] = au.Name
	claims["uuid"] = au.Uuid

	jwtPure := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := jwtPure.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", error_factory.NewInternalServerError("error generating token")
	}
	au.Jwt = jwt

	return jwt, nil
}

func GetNickNameFromJWT(tokens string) (string, *error_factory.RestErr) {
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
	nickName := claims["nick_name"]

	return (nickName).(string), nil
}
