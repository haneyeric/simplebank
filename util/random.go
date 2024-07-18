package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const letters = "qwertyuiopasdfghjklzxcvbnm"

// Retrun random int between min-max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Return random string from letters of length n
func RandomString(n int) string {
	var s strings.Builder
	l := len(letters)

	for i := 0; i < n; i++ {
		ch := letters[rand.Intn(l)]
		s.WriteByte(ch)
	}

	return s.String()
}

// Generate random owner
func RandomOwner() string {
	return RandomString(8)
}

// Generate random amount of money
func RandomBalance() int64 {
	return RandomInt(100, 1000000)
}

// Generate random currency
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD, JPY, CNY}
	l := len(currencies)
	return currencies[rand.Intn(l)]
}

// Generate random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
