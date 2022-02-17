package util

import (
	"fmt"
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

// RandomMoney returns a random money amount between 0 and 1000.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency from the supported currencies list.
func RandomCurrency() string {
	return supportedCurrencies[rand.Intn(len(supportedCurrencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@test.com", RandomString(6))
}
