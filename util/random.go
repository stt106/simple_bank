package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) // make sure it doesn't genereate duplicate random data.
}

// RandomInt generates a random int64 between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := rand.Intn(k)
		sb.WriteByte(alphabet[c])
	}
	return sb.String()
}

// RandomOwner generates a random owner of length 6
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	cs := []string{"EUR", "RMB", "USD", "GBP"}
	return cs[rand.Intn(len(cs))]
}
