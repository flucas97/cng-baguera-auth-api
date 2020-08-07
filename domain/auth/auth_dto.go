package auth

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	uuid "github.com/satori/go.uuid"
)

type auth struct {
	Name string `json:"name"`
	Uuid string `json:"_uuid"`
}

func New(name string) *auth {
	au := &auth{}
	au.Uuid = uuid.NewV4().String()
	au.Name = name
	return au
}

func (au *auth) GenerateJWT() (string, *error_factory.RestErr) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_name"] = au.Name
	claims["uuid"] = au.Uuid

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", error_factory.NewInternalServerError("error generating token")
	}

	return newToken, nil
}

func (a *auth) VerifyToken() bool {

	return true
}

func GetJWT(reqToken string, claims jwt.MapClaims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		fmt.Println("error")
	}
	return token, nil
}

func ValidateJWT(jwt *jwt.Token) string {
	for key, val := range jwt.Claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
