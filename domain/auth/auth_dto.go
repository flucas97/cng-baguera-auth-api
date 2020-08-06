package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Name string `json:"name"`
	Uuid uint64 `json:"_uuid"`
}

func CreateToken(authD Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_name"] = authD.Name
	claims["uuid"] = authD.Uuid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func (a *Auth) VerifyToken() bool {

	return true
}
