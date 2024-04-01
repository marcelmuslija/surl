package crypto

import (
	"crypto/rand"

	"github.com/btcsuite/btcutil/base58"
)

const shortCodeLength = 8

func GenerateShortCode() (string, error) {
	randomBytes := make([]byte, shortCodeLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	shortCode := base58.Encode(randomBytes)

	return shortCode[:shortCodeLength], nil
}
