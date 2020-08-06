package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type auth struct {
	Name string `json:"name"`
	Uuid string `json:"_uuid"`
}

func New(name string) (*auth, error) {
	au := &auth{}
	au.Uuid = uuid.NewV4().String()
	au.Name = name
	return au, nil
}

func (au *auth) GenerateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_name"] = au.Name
	claims["uuid"] = au.Uuid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func (a *auth) VerifyToken() bool {

	return true
}
