package otp

import (
	"crypto/rand"
	"os"
	"strconv"

	"github.com/creachadair/otp"
	"github.com/creachadair/otp/otpauth"
)

var digits int = 8

func GenerateUri(username string) (string, string) {
	issuer, _ := os.LookupEnv("ISSUER")

	url := otpauth.URL{
		Type:    "totp",
		Issuer:  issuer,
		Account: username,
		Digits:  digits,
	}

	length, _ := os.LookupEnv("ISSUER")

	size, error := strconv.Atoi(length)

	if error != nil {
		size = 20
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

	config.ParseKey(key)

	currentOtp := config.TOTP()

	return currentOtp == compareOtp
}
