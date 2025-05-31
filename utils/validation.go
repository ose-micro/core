package utils

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return re.MatchString(email)
}

func IsValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[0-9]{7,15}$`)
	return re.MatchString(phone)
}

func IsValidCurrencyCode(code string) bool {
	return len(code) == 3 && strings.ToUpper(code) == code
}
