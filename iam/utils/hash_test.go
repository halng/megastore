package utils

import (
	"github.com/stretchr/testify/assert"
	utils2 "github.com/tanhaok/MyStore/utils"
	"testing"
)

func TestComputeHMAC256(t *testing.T) {
	username := "changeme"
	email := "changeme@gmail.com"

	// Test HMAC
	hmac := utils2.ComputeHMAC256(username, email)
	if hmac == "" {
		t.Errorf("HMAC is empty")
	}
	assert.Equal(t, 44, len(hmac))
	assert.Equal(t, "/AfpN+JydNKzaX5IiT/4M3OqWQ1Hsws+UAougZj4ZRQ=", hmac)
}

func TestComputeMD5(t *testing.T) {
	data := []string{"string_1"}
	expectedHash := "058eb6ea2bdcc79a6a7208783c8bfb50"
	assert.Equal(t, expectedHash, utils2.ComputeMD5(data))
}
