package helper

import (
	"math/rand"
	"time"
)

// Function to generate random string
func GenerateRandomString(length int) string {
	// Characters to be used for generating random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice of the given length
	randomString := make([]byte, length)

	// Generate random string
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}
