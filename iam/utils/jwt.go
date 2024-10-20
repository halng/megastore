package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var (
	IdClaimKey       = "_id"
	UsernameClaimKey = "_username"
	RoleClaimKey     = "_role"
	EnvApiSecretKey  = "API_SECRET"
	TokenRequestKey  = "API"
)

func GenerateJWT(id string, username string, role string) (string, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)

	claims := jwt.MapClaims{}
	claims[IdClaimKey] = id
	claims[UsernameClaimKey] = username
	claims[RoleClaimKey] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))
}

// ExtractDataFromToken get data from token
func ExtractDataFromToken(tokenStr string) (bool, string, string, string) {
	/**
	*	token: uuid use this uuid to get actual token in cache, if exist => token valid, if not, token expire
	 */
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(EnvApiSecretKey)), nil
	})
	if err != nil {
		return false, "", "", ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims[IdClaimKey] != nil && claims[UsernameClaimKey] != nil && claims[RoleClaimKey] != nil {
		return true, claims[IdClaimKey].(string), claims[UsernameClaimKey].(string), claims[RoleClaimKey].(string)
	}
	return false, "", "", ""
}
