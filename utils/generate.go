package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateIdentity(prefix string) string {
	// Timestamp: YYYYMMDD-HHMMSS
	timestamp := time.Now().Format("20060102-150405")

	// Generate a 4-character alphanumeric tracking code
	trackingCode := GenerateCode(4)

	// Combine timestamp and tracking code
	return fmt.Sprintf("%s-%s-%s", prefix, timestamp, trackingCode)
}

// Generate a random alphanumeric tracking code
func GenerateCode(length int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	for i := range b {
		b[i] = letters[b[i]%byte(len(letters))]
	}

	return string(b)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateSlug(value string) string {
	value = strings.Map(func(r rune) rune {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ", r) {
			return r
		}
		return -1
	}, value)
	return strings.ToLower(strings.ReplaceAll(value, " ", "-"))
}

func GenerateKey(prefix string, length int) (string, error) {
	byteLength := length * 3 / 4

	// Generate random bytes
	b := make([]byte, byteLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Encode in URL-safe base64 and strip padding
	encoded := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)

	// Ensure it's exactly the desired length (trim if needed)
	if len(encoded) > length {
		encoded = encoded[:length]
	}

	return prefix + encoded, nil
}
