package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeHMAC256(t *testing.T) {
	username := "changeme"
	email := "changeme@gmail.com"

	// Test HMAC
	hmac := ComputeHMAC256(username, email)
	if hmac == "" {
		t.Errorf("HMAC is empty")
	}
	assert.Equal(t, 44, len(hmac))
	assert.Equal(t, "/AfpN+JydNKzaX5IiT/4M3OqWQ1Hsws+UAougZj4ZRQ=", hmac)
}
