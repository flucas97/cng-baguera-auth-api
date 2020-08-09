package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	uuid "github.com/satori/go.uuid"
)

type Token struct {
	Name     string `json:"name"`
	Uuid     string `json:"_uuid"`
	JwtToken string `json:"jwt"`
}

func New(name string) *Token {
	au := &Token{}
	au.Uuid = uuid.NewV4().String()
	au.Name = name
	return au
}

func (au *Token) GenerateJWT() *error_factory.RestErr {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nick_name"] = au.Name
	claims["uuid"] = au.Uuid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return error_factory.NewInternalServerError("error generating token")
	}

	au.JwtToken = newToken
	return nil
}

func ValidateJWT(reqToken string, claims jwt.MapClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		fmt.Println("error")
	}
	return token, nil
}
