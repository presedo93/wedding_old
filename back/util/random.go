package util

import (
	"fmt"
	"math/rand"
)

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	bytes := make([]byte, n)

	for i := 0; i < n; i++ {
		bytes[i] = byte(RandomInt(97, 122))
	}

	return string(bytes)
}

// RandomOwner generates a random owner name
func RandomName() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomPhoneNumber generates a random phone number
func RandomPhoneNumber() string {
	return fmt.Sprintf("+%d", RandomInt(100000000, 700000000))
}

// RandomID generates a 12 character random ID
func RandomID() string {
	return RandomString(12)
}
