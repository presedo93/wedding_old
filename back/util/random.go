package util

import "math/rand"

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
