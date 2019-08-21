package webtoken

import (
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	IsAdmin  int    `json:"is_admin"`
	IsRoot   int    `json:"is_root"`
	jwt.StandardClaims
}

var jwtSecret []byte

func ParseToken(token string) (*Claims, int) {
	code := e.Success

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, code
		}
	}
	if utils.CheckError(err, err.(*jwt.ValidationError).Errors) {
		code = e.Unauthorized
	}

	return nil, code
}

func GenerateToken(claims Claims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)

}
