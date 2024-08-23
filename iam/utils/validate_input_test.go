package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"gte=18,lte=65"`
}

// TestValidateInputTest tests the ValidateInput function
func TestValidateInputTest(t *testing.T) {
	t.Run("Valid input", func(t *testing.T) {
		input := TestStruct{
			Email:    "test@example.com",
			Username: "testuser",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.True(t, valid)
		assert.Nil(t, errors)
	})

	t.Run("Missing required field", func(t *testing.T) {
		input := TestStruct{
			Email:    "test@example.com",
			Username: "",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.Equal(t, "The username field is required", errors["0"])
	})

	t.Run("Invalid email format", func(t *testing.T) {
		input := TestStruct{
			Email:    "invalid-email",
			Username: "testuser",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.Equal(t, "The email field must be a valid email address", errors["0"])
	})

	t.Run("Field value less than minimum", func(t *testing.T) {
		input := TestStruct{
			Email:    "test@example.com",
			Username: "testuser",
			Age:      17,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.Equal(t, "The age field must be greater than or equal to 18", errors["0"])
	})

	t.Run("Field value greater than maximum", func(t *testing.T) {
		input := TestStruct{
			Email:    "test@example.com",
			Username: "testuser",
			Age:      66,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.Equal(t, "The age field must be less than or equal to 65", errors["0"])
	})

	t.Run("Missing multiple required fields", func(t *testing.T) {
		input := TestStruct{
			Email:    "",
			Username: "",
			Age:      30,
		}
		valid, errors := ValidateInput(input)
		assert.False(t, valid)
		assert.Equal(t, "The email field is required", errors["0"])
		assert.Equal(t, "The username field is required", errors["1"])
	})
}
