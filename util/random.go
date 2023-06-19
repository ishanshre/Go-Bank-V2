package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	// returns a random integer between min and max
	return min + rand.Int63n(max-min+1) // min -> max-min+min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	// generates and return random owner
	return RandomString(6)
}

func RandomMoney() int64 {
	// generates and returns a random money
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	// return a random currency
	currencies := []string{"EUR", "USD", "NPR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
