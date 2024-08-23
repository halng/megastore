package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
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

}
