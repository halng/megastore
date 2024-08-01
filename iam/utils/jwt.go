package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/tanhaok/MyStore/models"
	"os"
)

var (
	IdClaimKey       = "_id"
	UsernameClaimKey = "_username"
	RoleClaimKey     = "_role"
	EnvApiSecretKey  = "API_SECRET"
	TokenRequestKey  = "API"
)

func GenerateJWT(account models.Account) (string, error) {
	apiSecret := os.Getenv(EnvApiSecretKey)

	claims := jwt.MapClaims{}
	claims[IdClaimKey] = account.ID
	claims[UsernameClaimKey] = account.Username
	claims[RoleClaimKey] = "DEFAULT" // TODO: change this whenever mapping between account and role created

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))
}

// isValidToken will be renamed to getTokenFromCache
func isValidToken(tokenStr string) bool {
	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(EnvApiSecretKey)), nil
	})

	return err == nil
}

func ExtractTokenFromId(tokenStr string) (string, string, string) {
	/**
	*	token: uuid use this uuid to get actual token in cache, if exist => token valid, if not, token expire
	 */
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(EnvApiSecretKey)), nil
	})
	if err != nil {
		return "", "", ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%v", claims[IdClaimKey]), fmt.Sprintf("%v", claims[UsernameClaimKey]), fmt.Sprintf("%v", claims[RoleClaimKey])
	}
	return "", "", ""
}
