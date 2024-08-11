package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

var HMACSecretKey = "HMAC_SECRET"

func ComputeHMAC256(username string, email string) string {
	message := fmt.Sprintf("%s_%s", username, email)
	messageInByte := []byte(message)
	secret := os.Getenv(HMACSecretKey)
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write(messageInByte)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))

}
