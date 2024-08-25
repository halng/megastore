package utils

import (
	"fmt"
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

func GenerateJWT(id string, username string) (string, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)

	claims := jwt.MapClaims{}
	claims[IdClaimKey] = id
	claims[UsernameClaimKey] = username
	claims[RoleClaimKey] = "DEFAULT" // TODO: change this whenever mapping between account and role created

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
	if ok && token.Valid {
		return true, fmt.Sprintf("%v", claims[IdClaimKey]), fmt.Sprintf("%v", claims[UsernameClaimKey]), fmt.Sprintf("%v", claims[RoleClaimKey])
	}
	return false, "", "", ""
}
