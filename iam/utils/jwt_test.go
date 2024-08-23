package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	// Setup
	apiSecret := "testsecret"
	os.Setenv(EnvApiSecretKey, apiSecret)

	t.Run("Create and extract JWT", func(t *testing.T) {
		username := "changeme"
		id := "XXX-YYY-ZZZ"

		// Test JWT
		token, err := GenerateJWT(id, username)
		if err != nil {
			t.Errorf("Error generating JWT: %v", err)
		}

		if token == "" {
			t.Errorf("Token is empty")
		}

		// Test ExtractDataFromToken
		idFromToken, usernameFromToken, roleFromToken := ExtractDataFromToken(token)
		assert.Equal(t, id, idFromToken)
		assert.Equal(t, username, usernameFromToken)
		assert.Equal(t, "DEFAULT", roleFromToken)
	})
	t.Run("Extract data from invalid token", func(t *testing.T) {
		token := "invalid.bearer.token"
		idFromToken, usernameFromToken, roleFromToken := ExtractDataFromToken(token)
		assert.Equal(t, "", idFromToken)
		assert.Equal(t, "", usernameFromToken)
		assert.Equal(t, "", roleFromToken)
	})
	t.Run("Extract Token with missing claims", func(t *testing.T) {
		// Create a token with missing claims
		claims := jwt.MapClaims{
			IdClaimKey: "test-id",
			"exp":      time.Now().Add(time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, _ := token.SignedString([]byte(apiSecret))

		// Test ExtractDataFromToken
		id, username, role := ExtractDataFromToken(tokenStr)
		assert.Equal(t, "test-id", id)
		assert.Equal(t, "<nil>", username)
		assert.Equal(t, "<nil>", role)
	})
}
