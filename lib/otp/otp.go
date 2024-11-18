package otp

import (
	"crypto/rand"
	"os"
	"strconv"

	"github.com/creachadair/otp"
	"github.com/creachadair/otp/otpauth"
)

const digits int = 8

func GenerateUri(username string) (string, string) {
	issuer := "App"
	if value, exists := os.LookupEnv("ISSUER"); !exists {
		issuer = value
	}

	url := otpauth.URL{
		Type:    "totp",
		Issuer:  issuer,
		Account: username,
		Digits:  digits,
	}

	length := "20"
	if value, exists := os.LookupEnv("TOTP_SECRET_LENGTH"); !exists {
		length = value
	}

	size := 20
	if value, error := strconv.Atoi(length); error != nil {
		size = value
	}

	randomKey := make([]byte, size)

	rand.Read(randomKey)

	url.SetSecret(randomKey)

	return url.String(), url.RawSecret
}

func CompareOtp(key string, compareOtp string) bool {
	config := otp.Config{
		Digits: digits,
	}

	if config.ParseKey(key) != nil {
		return false
	}

	currentOtp := config.TOTP()

	return currentOtp == compareOtp
}
