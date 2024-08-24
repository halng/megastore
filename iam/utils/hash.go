package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

// ComputeMD5 to compute hash based on passed data
func ComputeMD5(data []string) string {
	var dataForHash string
	for _, s := range data {
		dataForHash += s
	}
	hash := md5.New()
	hash.Write([]byte(dataForHash))

	checkSum := hash.Sum(nil)
	return hex.EncodeToString(checkSum)

}
