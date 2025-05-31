package utils
import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecret() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	secret := base64.URLEncoding.EncodeToString(b)
	return secret
}